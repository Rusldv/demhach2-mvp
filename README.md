# demhack2-mvp (первый день)

Проект децентрализованного сервера, созданный на хакатоне DEMHACK-2

Что мы хотим:

1. Совместить технологии распределенных реестров и HTTP-протокола.
2. Обеспецить надежность хранения данных с использованием IPFS.
3. Сделать возможность идентификации приложений и личных аккаунтов по принципу ECDSA.

# demhack2-mvp (ночное)

В результате начной разработки мы реализовали:

1. Модель обработки HTTP-запросов на основе MVP-компонентов.
2. Обработку файла конфигурации для первоначальной инициальзации сервера.
3. Для добавления сайта достаточно создать дирректорию с именем сайта в дирректории root_dir (по умолчанию ./www).

О компонентах:
Компоненты позволяют отображать контент сайта, например, компонент главной страницы веб-сайта https://dorofeev.fun может загружаться и обрабатываться интерпретатором (github.com/traefik/yaegi/interp) из дирректории SRVDIR/www/dorofeev.fun/main/view.html. Дополнительно в дирректорию сайта может быть помещен файл options.json, в котором задаются локальные настройки сайта. На странице сайта можно использовать конструкции go templates для отображения динамического содержимого.

# demhack2-mvp (второй день)

![alt-текст](https://github.com/Rusldv/demhach2-mvp/blob/main/demhack2.png "DEMHACK 2")
