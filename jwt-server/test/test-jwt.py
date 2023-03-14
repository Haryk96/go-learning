import requests

params = {
    "name": "repo-deb",
    "age": "mirrors",
}
json = {
    "name": "test",
    "password": "password-tests"
}
res = requests.post(
    "http://localhost:8080/jwt",
    params=params,
    json=json   
)

print (res.text)