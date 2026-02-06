import os
import logging
from pymongo import MongoClient

client = None
db = None


def connect_db():
    global client, db

    # Read environment variables
    MONGO_HOST = os.getenv("MONGO_HOST", "localhost")
    MONGO_PORT = os.getenv("MONGO_PORT", "27017")
    MONGO_DB = os.getenv("MONGO_DB", "kindergarten")
    MONGO_USER = os.getenv("MONGO_USER")
    MONGO_PASSWORD = os.getenv("MONGO_PASSWORD")
    MONGO_AUTH_SOURCE = os.getenv("MONGO_AUTH_SOURCE", "admin")
    mongodb_uri=f"mongodb://{MONGO_USER}:{MONGO_PASSWORD}@{MONGO_HOST}:{MONGO_PORT}/{MONGO_DB}?authSource={MONGO_AUTH_SOURCE}"


    # mongodb_uri = os.getenv("MONGODB_URI")
    # database_name = os.getenv("DATABASE_NAME", "kindergarten")

    if not mongodb_uri:
        raise RuntimeError("MONGODB_URI is not set")

    client = MongoClient(mongodb_uri, serverSelectionTimeoutMS=10000)
    db = client[MONGO_DB]

    # Test connection
    client.admin.command("ping")
    logging.info("Connected to MongoDB successfully!")


def get_collection(name):
    return db[name]
