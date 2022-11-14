# importing the required modules  
from json import loads  
from kafka import KafkaConsumer  
#from pymongo import MongoClient  


# generating the Kafka Consumer  
my_consumer = KafkaConsumer(  
    'test001',  
     bootstrap_servers = ['192.168.1.67:9092'],  
#     auto_offset_reset = 'earliest',  
     enable_auto_commit = True,  
     group_id = 'my-group',  
     value_deserializer = lambda x : loads(x.decode('utf-8'))  
     )



for message in my_consumer:  
    message = message.value  
    #collection.insert_one(message)  
    print(message)  
    
         