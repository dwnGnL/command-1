package main

type Student struct {
	ID         int    // primary key
	Fname      string //
	Age        int    //
	IsStudent  bool   //
	IsWorker   bool
	IsTeacher  bool
	Average    float32
	Experience int
}

/*
1.Реализация select columns
a) Если колонки не существуют, вернуть ошибку неверно указано имя колонки(имя)
б) Реализовать оператор Where (<> ==)
2. Реализация Insert () into students
3. Реализация Update where id
4. Реализация Delete where id
5. Реалиция икремента
6. изменить метод create table tableName
7. Разделить программу на модули *
*/
