I’ve been mulling many of these points with respect to my projects and here are my thoughts:

1.  Fat models are okay at the start, but seem to become a problem the more the app grows
2.  A separate module for “business logic” is probably a good idea
3.  I think models are a “low level” interface to your data(base). And should not be much more.
4.  I call the business logic module “logic”. Others seem to call it “services”.
5.  When you need to implement a process that writes to one model then I think a manager is the way to go.
6.  When you need to implement a process that needs to write to two or more models, then I think you might want to use the “logic” module
7.  If you want to read any data that depends only on one model, then you can use model methods
8.  If you want to read data that depends on more than one model, you might want to use the “logic” module. The exception to this I think would be when a model is conceptually “nested” in another (via a  
    foreign key) then it may be natural to use model methods.
9.  The logic can be reused in views, api calls, celery tasks / cron jobs, management commands
10.  The logic module is a natural interface to your app.
11.  Forms are also a natural interface to your app. However I would keep form logic simple, and if it gets complicated (e.g. requires writing to two or more models) then factor that logic out into your logic module.

[Weblink](https://forum.djangoproject.com/t/where-to-put-business-logic-in-django/282/11)