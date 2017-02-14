This is an attempt at creating a service capable of controlling the heat
inside a room.

Author: Sylvain Prost

# How it work

A MQTT client register to the topic `/readings/motion` where he is notified
if there is activity in the room.

In such case he will register to `/readings/temperature` and by operating
savant computing publish the appropriate heat power needed to heat the room on
`/actuators/room-1`. This will work only if the sensor publish the temperature
at a frequency high enough. (other solution discussed below)

Otherwise if there is no activity it will simply close the heat power
and unsubscribe to the temperature topic.

With this behavior you the heating system will always be at 100% everytime there is
activity again or the service is restarted. This can be prevented by saving
sensors activity in a database. (solution is discussed later)

# Setup

Install go, at least 1.7. You can get it from [here](https://golang.org/)

Get a MQTT broker !

The simple way, if you already have docker:
 ```
    docker run -it -p 1883:1883 --name=mosquitto  toke/mosquitto
   ```

Extract the archive and run
```
    go run *.go -uri=your_mqtt_broker_u
   ```

# Possible Upgrade

## Valve power

For now the valve power is calculated with the brain of a middle school
student.
For better calculation you will want to know:
* How much heat is produced at each power level.
* How long does it take to eat the room.

More and better algorithm can be produced by including a efficient way to store
sensor data and actuator actions. They can then be collected and exploited
by expert and ai.

## THE DB

Building an IoT service with a db even for small project can greatly improve
its use cases and future upgrade.

It allow the logging of every sensor activity on witch you can apply
complex algorithm. In the case of our project it's nearly mandatory if
you want a innovative & smart heating system.

examples:

* Create an ai that can preview the usage and preheat room by looking at
the activity sensors. (maybe even include a calendar integration)
* How long does it take to heat a room ?
* Allow the profile by room. In case of multiple room you can parameter
which sensor/actuator belong to which room. Which also make a great tool to
see the impact of heating a room to other room.

## Remote API

Building this as a micro service and include an API.

## Mutiply sensors / integration

The more the better. Never enough data.

examples:

* Windows sensor, shut down heating system when open.
* Connected electricity. Know exactly the impact your heating system on your
bills. Can also preview energy spike and reduce the heat power.
* The weather

There is actually too much possibility...and it's starting to be late so
I hope you enjoyed reading this and looking through the code.

##### THE END


