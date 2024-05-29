Feature: Create object
   In order to manage system models
   I need to be able to create
   with objects using the repository
   Scenario: Create example object
      When i create a new user with the following details:
	     | Name 	| Link 		    |
	     | example  | examplelink1  |
      Then i should be able to get it by key

