using System.Threading.Tasks;
using Bank.Business.Entities;

namespace Bank.Business.Interfaces
{
    public interface ICustomerRepository
    {
        Task<Customer> AddAsync(Customer customer);
        Task<Customer> GetByDocumentAsync(string document);
    }
}