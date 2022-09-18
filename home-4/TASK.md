# Практическое задание
1. Добавить в пример с файловым сервером возможность получить список всех файлов
   на сервере (имя, расширение, размер в байтах)
2. С помощью query-параметра, реализовать фильтрацию выводимого списка по
   расширению (то есть, выводить только .png файлы, или только .jpeg)
3. *Текущая реализация сервера не позволяет хранить несколько файлов с одинаковым
   названием (т.к. они будут храниться в одной директории на диске). Подумайте, как
   можно обойти это ограничение?
4. К коду, написанному в рамках заданий 1-3, добавьте тесты с использованием
   библиотеки httptest.

# Примеры использования

Примеры предполагают, что приложение запускается из директории в которой находится директория "./files"

1) Запускаем приложение. В браузере по адресу "http://localhost:8082" можно наблюдать что-то типа:
<pre>
<table style="border: 1px solid black;"><tr><th style="border: 1px solid black;">name</th><th style="border: 1px solid black;">ext</th><th style="border: 1px solid black;">size</th></tr><tr><td  style="border: 1px solid black;">gopher</td><td style="border: 1px solid black;">png</td><td style="border: 1px solid black;">37023</td></tr><tr><td  style="border: 1px solid black;">outer</td><td style="border: 1px solid black;">txt</td><td style="border: 1px solid black;">40</td></tr><tr><td  style="border: 1px solid black;">testfile</td><td style="border: 1px solid black;">txt</td><td style="border: 1px solid black;">77</td></tr></table></pre>

2) Если вызвать приложение таким образом: "http://localhost:8082?ext=png", в таблице останутся файлы, только с расширением "png"
3) Файл index.html (можно открыть в браузере или запустить с помощью стороннего сервера) позволяет выбрать файл и загрузить его в директорию "/files"
4) Файл server_test.go позволяет тестировать функцию загрузки файла

## Примечание:
Тесты отображения файлов умею делать с помощью "Cypress", но это далеко за рамками этого курса.