""" Functions used for Currency Exchange  to find the following 

    Task 1 Estimate Value after exchange 
    Task 2 Currency left after an exchange 
    Task 3 Calculate Value of Bills 
    Task 4 Calculate numbers of bills 
    Task 5 Calculate leftover after exchanging into bills
    Task 6 Calculate value after exchange

"""


def exchange_money(budget, exchange_rate):
    """ Estimate Value after exchange
    :param budget: float - amount of money you are planning to exchange.
    :param exchange_rate: float - unit value of the foreign currency.
    :return: float - exchanged value of the foreign currency you can receive.
     
    """
    return budget / exchange_rate

def get_change(budget, exchanging_value):
    """
    :param budget: float - amount of money you own.
    :param exchanging_value: float - amount of your money you want to exchange now.
    :return: float - amount left of your starting currency after exchanging.
    """
    return budget - exchanging_value   
    
def get_value_of_bills(denomination, number_of_bills):
    """

    :param denomination: int - the value of a bill.
    :param number_of_bills: int - amount of bills you received.
    :return: int - total value of bills you now have.
    """
    return number_of_bills * denomination


def get_number_of_bills(budget, denomination):
    """

    :param budget: float - the amount of money you are planning to exchange.
    :param denomination: int - the value of a single bill.
    :return: int - number of bills after exchanging all your money.
    """
    return int(budget / denomination)    

def get_leftover_of_bills(budget, denomination):
    """

    :param budget: float - the amount of money you are planning to exchange.
    :param denomination: int - the value of a single bill.
    :return: float - the leftover amount that cannot be exchanged given the current denomination.
    """
    return budget % denomination
 
 
def exchangeable_value(budget, exchange_rate, spread, denomination):
    """

    :param budget: float - the amount of your money you are planning to exchange.
    :param exchange_rate: float - the unit value of the foreign currency.
    :param spread: int - percentage that is taken as an exchange fee.
    :param denomination: int - the value of a single bill.
    :return: int - maximum value you can get.
    """
    exchange_fee = (exchange_rate * (1 + spread/100 ))
    max_value_money= exchange_money(budget, exchange_fee)
    return int(max_value_money - get_leftover_of_bills(max_value_money, denomination))

    

