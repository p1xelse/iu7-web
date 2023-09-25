# iu7-web
## Цель работы, решаемая проблема/предоставляемая возможность

Write&Send - это онлайн платформа для обмена сообщениями, новостями и картинками между пользователями. 
Он позволяет в пределах одной страницы следить за новостями, которые публикуются пользователями, делиться ими, комментировать оценивать и сохранять в закладки к себе в профиль, коммуницировать как в личных сообщениях так и в беседах пользователям нашего сервиса. 

Личный блог, новостное онлайн сообщество, мессенджер с поддержкой вложений, поиск записей - проблемы которые решает наш продукт.


## Краткий перечень функциональных требований

1. Авторизация:
- Пользователю требуется ввести свой логин и пароль для успешной авторизации.
- При неверном вводе данных пользователь должен получить соответствующее уведомление о неправильном логине или пароле.
- Авторизация должна сохраняться для последующих сеансов использования приложения.

2. Регистрация:
- При регистрации пользователь должен указать уникальный логин и пароль.
- При вводе уже существующего логина пользователь должен получить ошибку о том, что данный логин занят.
- В процессе регистрации пользователь может заполнить дополнительные данные, такие как имя, фамилию, возраст и т.д.

3. Изменение профиля:
- Пользователь может изменять свои данные в профиле, такие как имя, фамилию, возраст и т.д.
- Изменения данных в профиле должны быть сохранены и отображаться после обновления страницы.

4. Написать пост:
- Пользователь может создавать новые посты, указывая заголовок, текст и прикреплять изображения или другие медиафайлы.
- Созданные пользователем посты должны отображаться на главной странице приложения.

5. Просмотр постов:
- Пользователь должен иметь возможность просматривать все посты, созданные другими пользователями на главной странице.
- Просмотр постов должен быть ограничен только авторизованными пользователями.

6. Изменение постов:
- Пользователь может изменять свои собственные посты, изменяя заголовок, текст и прикрепляя/удаляя изображения или другие медиафайлы.
- Изменения в постах должны быть сохранены и отображаться после обновления страницы.

7. Удаление постов:
- Пользователь может удалить свои собственные посты.
- После удаления поста, он должен быть недоступен для просмотра другим пользователям.

8. Поставить лайки:
- Пользователь может ставить лайки к постам других пользователей.
- Пользователь должен иметь возможность видеть общее количество лайков, полученных каждым постом.

9. Оставить комментарии:
- Пользователь может оставлять комментарии к постам других пользователей.
- Комментарии должны быть видны всем пользователям, просматривающим данный пост.

10. Добавить друзей:
- Пользователь может добавлять других пользователей в список друзей.
- Пользователь должен иметь возможность удалять друзей из своего списка.

11. Обмен сообщениями:
- Пользователь должен иметь возможность отправлять и получать сообщения от других пользователей.
- Сообщения должны быть видны только отправителю и получателю.
- Пользователь должен иметь возможность видеть историю сообщений с каждым контактом.

12. Создание сообществ:
- Пользователь может создавать свои собственные сообщества, указывая название, описание и изображение.
- Созданные сообщества должны отображаться на главной странице и быть доступными для просмотра другим пользователям.

13. Просмотр сообществ:
- Пользователь должен иметь возможность просматривать список всех существующих сообществ.
- В списке сообществ должна отображаться информация о названии, количестве участников и описании каждого сообщества.

14. Изменение сообществ:
- Пользователь может изменять название, описание и изображение своих собственных сообществ.
- Изменения в сообществе должны быть сохранены и отображаться после обновления страницы.

15. Удаление сообществ:
- Пользователь может удалить свои собственные сообщества.
- После удаления сообщества, оно должно быть недоступно для просмотра другим пользователям.


## Use-case диаграмма системы

![iu7-web-usecase](https://github.com/p1xelse/iu7-web/assets/78589385/04a3c033-dc3a-4a61-b7ae-1a5fa588eb3f)


## BPMN диаграмма основных бизнес-процессов

![iu7-web-bpmn](https://github.com/p1xelse/iu7-web/assets/78589385/03ecc0f8-b646-4734-8be5-772c79bcc79b)


## Примеры описания основных пользовательских сценариев

Пользователь может:
- Авторизоваться в сервсие
- Зарегистрироваться в сервисе
- Изменить данные в своем профиле
- Написать пост (создание, просмотр, изменение, удаление)
- Ставить лайки к постам
- Оставлять комментарии к постам
- Добавить друзей (добавление, удаление)
- Обмениваться сообщениями
- Создавать сообщества (создание, просмотр, изменение, удаление)

## ER-диаграмма сущностей

![](./assets/er.png)

## Диаграмма БД

![iu7-web](https://github.com/p1xelse/iu7-web/assets/78589385/43d1a431-1821-4504-b544-806151db485f)

## Компонентная диаграмма системы

Backend приложение построено на основе Clean architecture, что предполагает разделение приложения на три слоя:
- Repository - взаимодействие с данными(хранение, создание, получение,- изменение);
- Usecase - бизнес логика;
- Delivery - обработка запросов, отправка ответов.
![](./assets/component.png)

## Экраны будущего web-приложения на уровне черновых эскизов. Задача данного упражнения - понять, как с приложением должен взаимодействовать пользовать для упрощения проектирования API. Это могут быть классические wireframes, черновики от руки, наброски в PAINT/псевдографике/Figma.

Ссылочка на Figma с макетами: https://www.figma.com/file/NfEB0IRzh4rGnDJCyw796L/WS_service?type=design&node-id=0-1&mode=design


