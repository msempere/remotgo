# Remotgo

Send commands over ssh to AWS EC2 instances

### Example
Execute "df -H" command in all instances with highway as role.

![Usage example](http://i.imgur.com/69fyPJx.gif)

### Configuration

The following environment variables have to be properly set:
 - AWS_ACCESS_KEY_ID
 - AWS_SECRET_ACCESS_KEY
 - AWS_REGION
 
To use environment variables, do the following:

    $ export AWS_ACCESS_KEY_ID=<access_key>
    $ export AWS_SECRET_ACCESS_KEY=<secret_key>
    $ export AWS_REGION=<region>





### License
Distributed under MIT license. See `LICENSE` for more information.

