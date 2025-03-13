# Deliverables
## Deliverable 1 ✔
1. ✔ (10%) ER diagram: Create the ER diagram that corresponds to the above description.  
2. ✔ (8%) Relational database schema: Create the relational database schema that corresponds to your ER diagram. 
3. ✔ (7%) Integrity constraints: Define the necessary constraints that will ensure the correctness of the database to be created according to your relational database schema. These are primary keys, referential integrity constraints, domain and attribute constraints and user-defined constraints. Be inventive with the definition of user-defined constraints. 

## Deliverable 2
1. ✔ (10%) Database implementation: Implement the database according to your relational database schema and the constraints that you have defined.
    - Implemented in PR https://github.com/vpreseault/CSI2132-hotel-management/pull/16

2. (5%) Database population: Insert in your database data for each one of the 5 hotel chains. Each one of them has at least 8 hotels, which belong to at least 3 categories. Two of the hotels at least should be in the same area. Each hotel should have at least 5 rooms of different capacity. Populate your database with enough data to be able to showcase the execution of queries/triggers/views. 
    - [ ] 5 Chains
    - [ ] 8 Hotels for each chain (40 total)
        - [ ] 3 hotels in different categories
        - [ ] 2 hotels in same area/city
    - [ ] minimum 5 rooms for each hotel (200 total)
        - [ ] each room should have a different capacity
 
3. (10%) Database queries: Implement at least 4 queries of your choice on your database. Implement at least 1 query with aggregation and at least 1 with a nested query.  
    - [ ] 4 choice queries
        - [ ] Query X implmented in PR #Y
        - [ ] Query X implmented in PR #Y
        - [ ] Query X implmented in PR #Y
        - [ ] Query X implmented in PR #Y
    - 1 aggregation query
        - [ ] Query X implmented in PR #Y
    - 1 nested query
        - [ ] Query X implmented in PR #Y

4. (10%) Database modifications: Create the necessary SQL modifications (use queries and especially triggers): Your database should allow insert, delete and update operations of data in 
your database according to the referential integrity constraints, and moreover, to the user-defined constraints, which you have defined. Implement at least 2 triggers of your choice for this purpose. 
    - [ ] 2 triggers
        - [ ] Trigger X implmented in PR #Y
        - [ ] Trigger X implmented in PR #Y
    - 1 insert query
        - [ ] Query X implmented in PR #Y
    - 1 delete query
        - [ ] Query X implmented in PR #Y
    - 1 update query
        - [ ] Query X implmented in PR #Y

5. (5%) Database indexes: Implement at least 3 indexes on the relations of your database and justify why you have chosen these indexes: explain what type of queries and data updates you are expecting on your database and how these indexes are useful to accelerate querying of the database.
    - [ ] Index X implmented in PR #Y
    - [ ] Index X implmented in PR #Y
    - [ ] Index X implmented in PR #Y

6. (5%) Database views: Implement 2 views. View 1: the first view is the number of available rooms per area. View 2: the second view is the aggregated capacity of all the rooms of a specific hotel. You are welcome to implement more views of your choice. 
    - [ ] Number of available rooms implmented in PR #Y
    - [ ] Aggregated hotel room capacity implmented in PR #Y

7. (30%) Web application: Design and implement an appropriate User Interface, through which a user will be able to see the available rooms by giving different, multiple and combinations of criteria in order to choose the room that he/she is interested in and book it or rent it. These criteria should be: the dates (start, end) of booking or renting, the room capacity, the area, the hotel chain, the category of the hotel, the total number of rooms in the hotel, the price of the rooms. The user should be able to see the available choices when he/she changes the value of any of these criteria. 
    - [ ] Customer can search for rooms
    - [ ] Customer can book/rent room found in search
    - [ ] Search filters
        - [ ] start/end date 
        - [ ] room capacity
        - [ ] area/city
        - [ ] hotel chain
        - [ ] hotel category
        - [ ] total number of rooms in hotel
        - [ ] room price

The User Interface should allow the insert/delete/update of all information related to customers, employees, hotels and rooms. The user can be either a customer (who will use the interface in order to search for rooms and do bookings) or a hotel employee (who will use the interface to either turn a booking to renting when a customer checks in the hotel, or do directly a renting when a customer presents 
physically to the hotel). An employee should be able to insert a customer payment for a renting through the interface.
- [ ] insert/delete/update
    - [ ] customers
    - [ ] employees 
    - [ ] hotels
    - [ ] rooms 
- [ ] UI should allow employees to
    - [ ] transform bookings to rentings
    - [ ] create rentings for customers directly
    - [ ] insert customer payment info

The user should be able to see in the User Interface the two specific SQL Views implemented 
in (7/2f). 
- [ ] Display number of available rooms view
- [ ] Display aggregated hotel room capacity view

The User Interface should be user friendly, meaning that the user is not required to know SQL. All information should be presented to the user through appropriately designed forms. You should use appropriate elements, like drop-down lists, radio buttons etc. 

# Video
Create a video file .mp4 that is not more than 30MB and the time length of it is around 10 minutes and 
at most 15 minutes. You can use any software you want in order to create the video. The video should 
record your screen and your voice while you are presenting and describing your project. 
In the video you are asked to give the following information in the following order: 
1. Mention the software technologies that you have used for the creation of your project.  
2. Show briefly the relational database schema you have implemented and comment on any 
differences of this schema and the one you have submitted as part of Deliverable 1.  
3. Comment briefly on the major integrity constraints that you have implemented. Justify their 
choices and show the respective SQL code. 
4. Show and describe briefly the data with which you have populated your database. For 
example, mention how you created the data and what is the size in number of rows per major 
relation. 
5. Show the execution of some SQL queries on your database. 
6. Show the execution of some modifications on your database which fire a trigger. Show the 
SQL code of the trigger and explain how it implements a specific user-defined constraint. 
7. Show the SQL code and comment on the indexes you have implemented on your database. 
8. Show the SQL code and comment on the views you have implemented on your database. 
9. Show the User Interface you have implemented. Show briefly some of the functionalities of 
the interface: e.g. you can show the insertion/deletion/update of data, you can show the 
customer and the employee view and explain how they differ. Furthermore, show how a new 
booking or renting is made, and show how a booking is transformed to a renting when the 
customer checks in.
10. Fill out timestamp table
