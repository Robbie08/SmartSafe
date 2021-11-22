import requests as req
import random
from twilio.rest import Client

#Account_id and auth_tok is from twilio 
account_id = 'AC787cc50b1cfe9876d6b56b7bdef62699'
auth_tok = '6960bff9ba8ff13e2f864a0d89873798'
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
password = generatePassword(10)

#make Client to send message
client = Client(account_id,auth_tok)

#This sends the message to a given phone number
message = client.messages.create(
    body = password,#the password that is generated
    from_= '+14342660608',#twilio phone number
    to = '+16199882865' #This can be any phone number, must in this format +1##########
)
print(message.body)
sendPayload(password) # Generate password, package it into a payload and send it via http
