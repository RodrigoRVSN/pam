from flask import Flask, jsonify, request
import mysql.connector

app = Flask(__name__)


def get_connection():
    return mysql.connector.connect(
        host="localhost", user="root", password="password", database="task_management"
    )


@app.route("/users", methods=["GET"])
def get_users():
    connection = get_connection()
    cursor = connection.cursor()
    cursor.execute("SELECT * FROM Users")
    result = cursor.fetchall()
    return result


if __name__ == "__main__":
    app.run(debug=True)
