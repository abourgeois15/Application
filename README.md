# Go trainning, Full Stack Application

 This full stack application had been developped in the context of the Go training. The training is focused on the assimilation of the main concept of Go language used to develop the backend of a real application. The frontend is not mandatory and is then not very thorough.

 ## Description 

 This application is an item viewer based on the video game factorio. The idea is to be able to get easily information about items and machines from the game: Crafting time, type of machine, recipe, etc... The user can simply navigate between pages by clicking on elements displayed. A CRUD interface is available to Create, Read, Update and Delete any item.

 The application is composed by 3 elements: the api (backend), the client (frontend) and the database.
 * The api is developped with the go coding language.
 * The client is developped with React JS.
 * The database interface is MySQL.

 ## How to start the application

 The first step is to clone the git repository and open a terminal in the root of the project.
 Then enter the command ``docker-compose up``. This should create the 3 docker containers and start them.
 Finaly go to http://localhost:3000.