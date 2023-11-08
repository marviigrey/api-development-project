In this project we will be creating an API to handle all requests to our database.
We will be creating a mysql string connection to have permission into the database 
by passing the username,password and the name of the database. 

Step 1: configure and install mysql on your machine,create a database and add few
items inside:

    sudo apt update
    sudo apt install mysql-server
    #set security options:
    sudo mysql_secure_installation
    #login to mysql
    sudo mysql -u root -p
    //create a database called inventory
    CREATE DATABASE inventory;
    use inventory;
    create table products(
    id int NOT NULL AUTO_INCREMENT,
    name varchar(225) NOT NULL,
    quantity int,
    price float(10,7),
    PRIMARY KEY(id)
    );
    insert into products values(1, "chair", 100, 200.00);
    insert into products values(1, "desk", 300, 600.00);
    
We have successfully configured and setup or mysql server and configured
our database too.

Step 2: 
	create a working directory/repository for our project source code.
	after creating the directory we initialize a module inside the 
	directory:
	
			go mod init "example/api-dev"
			
Download the modules we will need for this project:
			go get github.com/gorilla/mux
			go get github.com/go-sql-driver/mysql
			
Create a constants.go file where you will declare your mysql database login details
as variables for us to point to them in our code.
			
create a file called app.go that will store information about routers,
DB variables and methods that will handle our routes.

inside our app.go:

-create a struct that will have two values, our router information and the
database information.

		type App struct {
			Router *mux.Router
			DB	*sql.DB
		}

	We move to creating two methods for our struct values now. The methods
	will have a pointer that will be pointig to our struct that we created
	earlier on. The first method will be a connection string that will be 
	opening a database connection.
- We create the method for SQL connectionString as Initialise(). This method
will contain 2 variables. The connection string that links to the
Database credentials defined in the constants.go file. and the app.DB
that allows an open connection to a mysql database using an SQL 
driver. Next we create mux router variable using the NewRouter function. 
we declared both our http Router and the database string connection inside
the iniitialise() method.
- Create the second method called run to pass the mux.router as our router
and listen on the right address.
- create the main.go file. The main.go file will contain the entry point to
  our functions.

  
    

    
    