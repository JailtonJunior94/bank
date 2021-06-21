using System.Threading.Tasks;
using Bank.Business.Entities;
using Bank.Business.Interfaces;

namespace Bank.Business.Repositories
{
    public class CustomerRepository : ICustomerRepository
    {
        public CustomerRepository()
        {

        }

        public Task<Customer> AddAsync(Customer customer)
        {
            throw new System.NotImplementedException();
        }

        public Task<Customer> GetByDocumentAsync(string document)
        {
            throw new System.NotImplementedException();
        }
    }
}
