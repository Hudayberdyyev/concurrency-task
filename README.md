# Тестовое задание на знание goroutine

У нас есть сущность менеджера (**manager**) который умеет выполнять задачи (**task**). Менежер должен уметь принимать задачи и выполнять их. Прошлый бэкенд разработчик который недавно уволился, написал уже готовый функционал. Но его программа работает недостаточно быстро. Ваша задача состоит в том чтобы улучшить программу так, чтобы она работала намного быстрее. </br>
Все изменения должны быть выполнены на файле manager.go, так чтобы тесты не сломались.</br>
Удачи !

**Условия выполнения задания: </br>**
* в корневой папке выполняем команду `go test -v -race core/*`
* после выполнения вы должны увидеть ошибку, о том что тесты не проходят. </br> 
(так как программа медленная)
* ваша задача изменить manager.go так чтобы этот тест прошёл.