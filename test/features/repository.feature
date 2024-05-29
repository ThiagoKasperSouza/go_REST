Feature: Create object
   In order to manage system models
   I need to be able to create
   with objects using the repository
   Scenario: Create example object
      Given my redis client is running
      When i create a new object with the following details:
	     | Name 	| Link 		    |
	     | example  | examplelink1  |
      Then i should be able to get it by key
   Scenario: List example objects
      Given my redis client is running
      Then i should be able to list objects from the example key
   Scenario: Update object
      Given my redis client is running
      And i have at least one object on the client
      Then i should be able to update this object
   Scenario: Delete object
      Given my redis client is running
      And i have at least one object on the client
      Then i should be able to delete this object


