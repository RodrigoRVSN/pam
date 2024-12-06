import mysql.connector

db = mysql.connector.connect(
    host="localhost", user="root", password="password", database="task_management"
)

cursor = db.cursor()

create_users_table = """
    CREATE TABLE IF NOT EXISTS Users (
        id INT AUTO_INCREMENT PRIMARY KEY,
        name VARCHAR(255) NOT NULL,
        email VARCHAR(255) NOT NULL,
        password VARCHAR(255) NOT NULL,
        created_at DATETIME 
    )
"""

create_tasks_table = """
    CREATE TABLE IF NOT EXISTS Tasks (
        id INT AUTO_INCREMENT PRIMARY KEY,
        title VARCHAR(255) NOT NULL,
        description VARCHAR(255) NOT NULL,
        user_id INT NOT NULL,
        due_date DATETIME,
        FOREIGN KEY(user_id) REFERENCES Users(id)
    )
"""

cursor.execute(create_users_table)
cursor.execute(create_tasks_table)
db.commit()

cursor.close()
db.close()

print("Tables creation done")
