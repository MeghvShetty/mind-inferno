
# Read-only operations (Serve car data, Filter Cars)
# Parameters (Query parameters , Path Parameters)
# Type hints (vaildateion,  conversion)

from fastapi import FastAPI, HTTPException
import uvicorn
from schema import load_db

apiserver = FastAPI()
db = load_db()

@apiserver.get("/")
async def welcome(name):
    """Returns a message for get request"""
    return {'message':"Welcome "+name+" to my car Sharing service!"}

@apiserver.get("/date")
async def date():
    """Returns a message for get request"""
    return {'date':"Sun 5 oct 2025"}


# Please add an operation called get_cars()
# That is served at /api/cars
# And that returns all car data 

@apiserver.get("/api/cars")
def get_cars(size:str| None = None, doors:int| None = None) -> list:
    if  size:
        return [car for car in db if car.size == size]
    if doors:
        return [car for car in db if car.doors >= doors] 
    return db

@apiserver.get("/api/cars/{id}")
def car_by_id(id: int) -> dict:
    result = [car for car in db if car.id == id]
    if result:
        return result[0]
    else:
        raise HTTPException(status_code=404, detail=f"Car not found with id={id}.")


if __name__ == "__main__":
    uvicorn.run("carsharing:apiserver",reload=True)