Go version 1.15.1

Run test in CoffeeMachineService_test to run test the machine

Note :
1) Followed YAGNI while writing this code
2) Could not do TDD because of time crunch but tried where ever it was possible and important
3) Tried to code in a way that this can be extensible
4) Haven't implemented parallel drink as we are running this through test only. Using go routines this can de 
done easily as I have added locks on domains updates! We can fire N go routine and call makeBeverage() to run in parallel

IMPORTANT NOTE : 
Haven't added unit test for Alert Event Publisher or for domains due to time crunch, but we can test the alert
publisher by mocking and write UT for domain.

