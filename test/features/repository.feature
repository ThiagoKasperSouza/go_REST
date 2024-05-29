Feature: Create object
   In order to manage system models
   As sys admin
   I need to be able to create
   with objects using the repository
   Scenario: Create example object
      Given i have admin permissions
      When i create a new user with the following details:
	     | Name 	| Link 		    |
	     | example  | examplelink1  |
      Then i should create this object sucessfully

