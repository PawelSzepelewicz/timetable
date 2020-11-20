# timetable

Telegram-бот для расписания.

Прежде чем запускать, необходимо получить токен в **BotFather**.
Что бы это сделать отправляем ему команду `/newbot`, выбираем имя, которое будет отображаться в списке контактов, и адрес.
Если адрес не занят, а имя введено правильно, **BotFather** пришлёт в ответ сообщение с токеном — «ключом» для доступа к созданному боту. Его нужно сохранить и никому не показывать.

Полученный токен необходимо добавить в `config.json`.

## Локальный запуск

1. Установить [GIT](https://git-scm.com/book/ru/v1/%D0%92%D0%B2%D0%B5%D0%B4%D0%B5%D0%BD%D0%B8%D0%B5-%D0%A3%D1%81%D1%82%D0%B0%D0%BD%D0%BE%D0%B2%D0%BA%D0%B0-Git#%D0%A3%D1%81%D1%82%D0%B0%D0%BD%D0%BE%D0%B2%D0%BA%D0%B0-%D0%B2-Windows).
2. Установить [Golang](https://golang.org/dl/).
3. Клонировать репозиторий - `git clone https://github.com/alxshelepenok/timetable.git`.
4. В дирректории с ботом выполнить команду - `go run main.go`.

## Развертывание в Heroku

1. Установить [GIT](https://git-scm.com/book/ru/v1/%D0%92%D0%B2%D0%B5%D0%B4%D0%B5%D0%BD%D0%B8%D0%B5-%D0%A3%D1%81%D1%82%D0%B0%D0%BD%D0%BE%D0%B2%D0%BA%D0%B0-Git#%D0%A3%D1%81%D1%82%D0%B0%D0%BD%D0%BE%D0%B2%D0%BA%D0%B0-%D0%B2-Windows).
2. Клонировать репозиторий - `git clone https://github.com/alxshelepenok/timetable.git`.
3. Создать аккаунт и установить [Heroku](https://www.heroku.com/).
4. Создать go-проект через интерфес и следовать дальнейшим инструкциям.

В настоящее время Heroku позволяет вам разместить ваше приложение бесплатно. Это отлично подходит для размещения вашего проекта в демонстрационных целях.
Тем не менее, если у вас бесплатный план, в соответствии с официальной документацией Heroku, приложение засыпает после 30 минут бездействия.
Что бы этого избежать, вы можете пинговать приложение каждые 30 минут с помощью скрипта в Google Sheet.

1. Создайте новый документ Google Sheet.
2. Введите `1` в ячейку `B1`, которая будет работать как счётчик
3. Перейдите в Инструменты -> Редактор скриптов.
4. Добавьте следующий код:

	```
	function myFunction() {
		 var d = new Date();
		 var timeZone = Session.getScriptTimeZone();
		 var hours = d.getHours();
		 var currentTime = d.toLocaleDateString();
		 var counter = SpreadsheetApp.getActiveSheet().getRange("B1").getValues();
		 if (hours >= 6 && hours <= 23) {
			 var response = UrlFetchApp.fetch("https://timetable.herokuapp.com/");
			 SpreadsheetApp.getActiveSheet().getRange("A" + counter).setValue("Visted at " + currentTime + " " + hours + "h" + " " + timeZone);
			 SpreadsheetApp.getActiveSheet().getRange("B1").setValue(Number(counter) + 1);
	 	}
	}
	```

5. Сохраните этот скрипт.
6. Добавьте новый триггер, выберите триггер по времени, минутный таймер и интервал, в течении которого этот скрипт должен посещать ваше приложение Heroku. И сохраните.

## Лицензия
The MIT License (MIT)

Copyright (c) 2019 Alexander Shelepenok

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
