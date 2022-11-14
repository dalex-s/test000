# importing the required libraries  
from time import sleep
from json import dumps
from kafka import KafkaProducer


# initializing the Kafka producer  
my_producer = KafkaProducer(  
    bootstrap_servers = ['192.168.1.67:9092'],  
    value_serializer = lambda x:dumps(x).encode('utf-8')  
    )
for n in range(1000):  
    #my_data = {'num' : n,'any' : n+n}  
    my_data = {'x':n}
    my_producer.send('test001', value = my_data)  
    sleep(0.01)  
    