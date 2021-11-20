import requests as req
import random

# This function will send our payload with password
# @params: A password that we will send over to our http server
def sendPayload(passwd):
    payload = {'message': passwd, 'id':'003349' }
    req.post('http://localhost:8080/pyclient', params=payload)

# This function will generate a password
# @params length: determines the number of characters in our password
# @return : generated password
def generatePassword(length):
    return ''.join([chr(random.randint(65,127)) for _ in range(length)])

sendPayload(generatePassword(10)) # Generate password, package it into a payload and send it via http
