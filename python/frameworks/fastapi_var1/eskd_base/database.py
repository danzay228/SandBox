from starlette.config import Config
from sqlalchemy import create_engine
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy.orm import sessionmaker

config = Config(".env")

SQLALCHEMY_DATABASE_URL = f"postgresql://{config('PGUSER')}:{config('PGPASSWORD')}@{config('PGHOST')}:{config('PGPORT')}/{config('PGDATABASE')}"

engine = create_engine(SQLALCHEMY_DATABASE_URL)
SessionLocal = sessionmaker(autocommit=False, autoflush=False, bind=engine)
Base = declarative_base()


def get_db():
    db = SessionLocal()
    try:
        yield db
    finally:
        db.close()
