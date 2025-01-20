Ссылка на копирование репозитория

git clone https://github.com/arikchakma/backend-projects.git cd backend-projects/task-tracker

Забилди, запускай

go build -o task-tracker ./task-tracker --help # To see the list of available commands
Чтоб добавить задачу

./task-tracker add "Buy groceries"
Чтоб обновить задачу

./task-tracker update 1 "Buy groceries and cook dinner"
Чтоб удалить задачу

./task-tracker delete 1
чтоб отметить задачу

./task-tracker mark-in-progress 1 ./task-tracker mark-done 1 ./task-tracker mark-todo 1
чтоб вывести весь список задач
./task-tracker list ./task-tracker list done ./task-tracker list to-do ./task-tracker list in-progress