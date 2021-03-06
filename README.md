# Yandex Practicum Go DevOps
Учебный проект на [Яндекс Практикуме](https://practicum.yandex.ru/). Содержит реализацию сервиса по сбору и хранению метрик.

## Структура проекта

```
├── .github
├── cmd
│   ├── agent               Агент собирает, буферизует и экспортирует метрики
│   └── server              Сервер предоставляет внешний интерфейс для приемки и агрегации метрик
├── internal
│   ├── commons             Shared компоненты, не относящиеся к конкретному пакету
│   │   ├── executor            Executor содержит общую функциональность, присущую сборщикам/экспортерам метрик
│   │   ├── handlers            Общая функциональность обработчиков запросов, например проставление HTTP-заголовков 
│   │   ├── logger              Функционал логирования, под капотом используется zerolog
│   │   ├── routing             Роутер, используемый сервером и в handler-ах
│   │   └── templating          Парсеры шаблонов, применяющие данные к каким-либо шаблонам
│   └── metrics             Пакет для работы с метриками, основной пакет в приложении
│       ├── buffering           Компоненты, реализующие буферизацию метрик перед экспортированием
│       ├── delivery            Обработчики запросов, которые использует сервер
│       ├── domain              Доменные модели метрик и константы, использующиеся и агентом, и сервером
│       ├── executors
│       │   ├── collectors          Сборщики метрик, которыми пользуется агент
│       │   └── exporters           Экспортеры метрик, которыми пользуется агент
│       ├── rendering           Компоненты, реализующие рендеринг метрик, например для отображения на HTML-страницах
│       ├── repository          Компоненты, реализующие хранение метрик, например в базе данных или памяти сервера
│       └── service             Компонент, реализующий основную бизнес-логику по работе с метриками,
│                                          предоставляет интерфейсы остальным компонентам
│
└── web
```

## Обновление шаблона

Чтобы получать обновления автотестов и других частей шаблона, выполните следующую команду:

```
git remote add -m main template https://github.com/yandex-praktikum/go-musthave-devops-tpl.git
```

Для обновления кода автотестов выполните команду:

```
git fetch template && git checkout template/main .github
```
