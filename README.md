# Что это?
Это репозиторий с примером кода из вебинара и домашним заданием для следующего урока.

# Как сдавать ДЗ?
Делаете форк репозитория, клонируете его к себе, выполняете дз и предлагаете пулл-реквест.  
В данном случае каждое задание можно реализовать в отдельном файле. Если есть желаение, можно связать
это всё в одну программу, которая будет спрашивать что нужно сделать (найти строку на сайтах или скачать файл) и
выполнять нужный метод.

# ДЗ
1) Напишите функцию, которая будет получать на вход строку с поисковым запросом (string) и массив ссылок на страницы, по которым стоит произвести поиск ([]string). Результатом работы функции должен быть массив строк со ссылками на страницы, на которых обнаружен поисковый запрос. Функция должна искать точное соответствие фразе в тексте ответа от сервера по каждой из ссылок. 
Поисковую строку и массив можете либо запросить у пользователя, либо захардкодить.

2) Напишите функцию, которая получает на вход публичную ссылку на файл с "Яндекс.Диска" и сохраняет полученный файл на диск пользователя.
Документация по теме: https://yandex.ru/dev/disk/api/reference/content-docpage/
**Это задание со звёздочкой. Выполнение не обязательно, но очень приветствуется**

### anonymous student
<p>

[findlinks](findlinks/findlinks.go), [findlinks_test](findlinks/findlinks_test.go) - Write a function that will receive an input string with a search query (string) and an array of links to pages that should be searched ([]string). The result of the function should be an array of strings with links to the pages on which the search query was found. The function should look for an exact match to the phrase in the response text from the server for each of the links.

</p>
<p>

[yandex](yandex/yandex.go), [yandex_test](yandex/yandex_test.go) - Write a function that receives a public link to the file from Yandex.Disk as an input and saves the received file to the user's disk.
</p>