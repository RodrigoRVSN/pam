import datetime
from time import time
from typing import TypedDict
from flask import Flask, Request, jsonify, request
import mysql.connector
from celery import Celery

celery = Celery("main", broker="pyamqp://guest:guest@172.17.0.1//", backend="rpc://")

app = Flask(__name__)


def get_connection():
    return mysql.connector.connect(
        host="localhost", user="root", password="password", database="task_management"
    )


@app.get("/users")
def get_users():
    connection = get_connection()
    cursor = connection.cursor()
    cursor.execute("SELECT * FROM Users")
    result = cursor.fetchall()
    return jsonify(result)


@app.get("/tasks")
def get_tasks():
    connection = get_connection()
    cursor = connection.cursor()
    cursor.execute("SELECT * FROM Tasks")
    result = cursor.fetchall()
    return jsonify(result)


class UserRequest(TypedDict):
    name: str
    email: str
    password: str


@app.post("/create-user")
def create_user():
    if not request.json:
        return jsonify({"error": "Invalid or missing JSON payload"}), 400
    data: UserRequest = request.json

    connection = get_connection()
    cursor = connection.cursor()
    try:
        query = """
            INSERT INTO Users (name, email, password, created_at)
            VALUES (%s, %s, %s, %s)
        """
        values = (
            data["name"],
            data["email"],
            data["password"],
            datetime.datetime.now(),
        )
        cursor.execute(query, values)
        connection.commit()
        return jsonify({"message": "User created successfully"}), 201
    except Exception as e:
        return jsonify({"error": str(e)}), 500
    finally:
        connection.close()


class TaskRequest(TypedDict):
    title: str
    description: str
    user_id: str
    due_date: str


@celery.task
def schedule_notification(task_title, user_id):
    print(f"Task {task_title} for user {user_id}")


@app.post("/create-task")
def create_task():
    if not request.json:
        return jsonify({"error": "Invalid or missing JSON payload"}), 400
    data: TaskRequest = request.json

    connection = get_connection()
    cursor = connection.cursor()
    query = """
        INSERT INTO Tasks (title, description, user_id, due_date)
        VALUES (%s, %s, %s, %s)
    """
    values = (data["title"], data["description"], data["user_id"], data["due_date"])

    due_date = datetime.datetime.strptime(data["due_date"], "%Y-%m-%d %H:%M:%S")
    lead_time = datetime.timedelta(minutes=10)
    notification_time = due_date - lead_time
    schedule_notification.apply_async(
        args=[data["title"], data["user_id"]], eta=notification_time
    )
    cursor.execute(query, values)
    connection.commit()
    return jsonify({"message": "Task created successfully"}), 201


if __name__ == "__main__":
    app.run(debug=True)
