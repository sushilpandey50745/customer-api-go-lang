Three subdirs
## domain-> library-> Consist of domain model "Customer" and persistence Abstraction
    Create struct "Customer"_> ID, Name, Email
    Create interface CustomerStore to specify CRUD-> Create(Customer)
    Update(string, Customer)error,
    Delete(string)error
    GetById(string)(Customer,error)
    GetAll()([]Customer,error)

## mapstore->library-> A persistance layer for domain model(customer). 
    store to implement CRUD operations of domain.Customer
    type MapStore struct{

    }

## cmd(main)
    Add a Customer in Store
    Get Customer Details
    List Customers
    Update Customer
    List Customer
    Delete Customer
    List Customer