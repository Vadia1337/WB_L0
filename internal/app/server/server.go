package server

import (
	"WB_L0/internal/app/cache"
	"WB_L0/internal/app/model"
	"WB_L0/internal/app/store/pgstore"
	"context"
	"database/sql"
	"fmt"
	"github.com/nats-io/stan.go"
	"html/template"
	"log"
	"net/http"
	"time"
)

type ViewData struct {
	Title string
}

type Server struct {
	config *Config
	store  *pgstore.Store
	nats   stan.Conn
	cache  *cache.Cache
}

func New(config *Config) *Server {
	return &Server{
		config: config,
	}
}
func (s *Server) Start() {
	s.newDB()

	s.tryConnectToRedis()

	s.tryConnectToNats()

	s.initWebServer()
}

func (s *Server) newDB() {
	db, err := sql.Open("postgres", s.config.DbUrl)
	if err != nil {
		log.Print(err)
	}

	if err := db.Ping(); err != nil {
		log.Print(err)
	}

	s.store = pgstore.New(db)
}

func (s *Server) createOrder(order *model.Order) {

	//проверить, есть ли такой ордер в бд
	err := s.store.Order().GetCount(order.OrderUid)
	if err == nil {
		log.Print("такой ордер уже есть в бд!")
		return
	}

	err = s.store.Delivery().Create(&order.Delivery)
	if err != nil {
		log.Fatal(err)
	}

	err = s.store.Payment().Create(&order.Payment)
	if err != nil {
		log.Fatal(err)
	}

	err = s.store.Order().Create(order)
	if err != nil {
		log.Fatal(err)
	}

	for idx, _ := range order.Items {
		err = s.store.Item().Create(&order.Items[idx])
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Сохранили order в бд")
}

func (s *Server) connectToNats() (stan.Conn, error) {
	nats, err := stan.Connect(s.config.ClusterID, s.config.ClientID, stan.NatsURL(s.config.NatsURL),
		stan.Pings(10, 2),
		stan.SetConnectionLostHandler(func(_ stan.Conn, reason error) {
			fmt.Println("Потерял соединение, пытаюсь восстановить...")
			s.tryConnectToNats()
			return
		}))
	if err != nil {
		return nil, err
	}

	fmt.Println("Присоединился к nats")

	return nats, nil
}

// Функция на случай, если сервис nats падает, наш сервис старается
// переподключиться в надежде, что nats поднимится
func (s *Server) tryConnectToNats() {
	nats, err := s.connectToNats()
	if err != nil {
		log.Print(err)
		s.tryConnectToNats()
	}
	s.nats = nats

	s.subscribe()

	duration := time.Second
	time.Sleep(duration)
}

// Подключение к редису и наполнение данными из бд
// Используется для первого и повторных подключений к редису
// в случае, если он упал
func (s *Server) tryConnectToRedis() {
	redis, err := cache.New(s.config.RedisIp)
	if err != nil {
		log.Print(err)
	}

	s.cache = redis

	orders, err := s.store.Order().GetAll()
	if err != nil {
		log.Print(err)
	}

	ctx := context.Background()
	for _, v := range orders {
		err = s.cache.Save(ctx, v.OrderUid, &v)
		if err != nil {
			log.Print(err)
		}
	}

}

func (s *Server) subscribe() {
	_, err := s.nats.Subscribe("foo", func(m *stan.Msg) {
		fmt.Println("flow from nats ...")
		s.receivedMsg(m.Data)
	}, stan.DurableName("my-durable"))
	if err != nil {
		log.Fatal(err)
	}
}

func (s *Server) receivedMsg(msg []byte) {
	order := model.Order{}
	isEquals := model.Equals(&order, msg)
	if isEquals == true {
		fmt.Println(order)
		err := s.cache.Save(context.Background(), order.OrderUid, &order)
		if err != nil {
			log.Print(err)
		}
		s.createOrder(&order)
	}
}

func (s *Server) initWebServer() {

	http.HandleFunc("/", s.HandleMain())

	http.HandleFunc("/order", s.HandleOrder())

	fmt.Println("Server is listening...")
	err := http.ListenAndServe(":8186", nil)
	if err != nil {
		log.Print(err)
	}
}

func (s *Server) HandleMain() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl, _ := template.ParseFiles("web/index.html")
		err := tmpl.Execute(w, nil)
		if err != nil {
			log.Print(err)
		}
	}
}

func (s *Server) HandleOrder() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		uid := r.PostFormValue("uid")
		order, err := s.getOrderByUid(uid)
		if err != nil {
			log.Print(err)
		}
		// отдаем в json для тестов, можно еще внутри метода cache.Get
		// убрать Unmarshal
		//orderJson, _ := json.Marshal(order)
		//w.Write(orderJson)

		tmpl, _ := template.ParseFiles("web/order.html")
		err = tmpl.Execute(w, order)
		if err != nil {
			log.Print(err)
		}
	}
}

func (s *Server) getOrderByUid(uid string) (*model.Order, error) {

	order, err := s.cache.Get(context.Background(), uid)
	if err != nil {
		log.Print(err)

		log.Print("В кэше нет - отдали из БД")
		return s.store.Order().GetOne(uid)
	}

	log.Print("отдали из кэша")
	return order, nil

}
