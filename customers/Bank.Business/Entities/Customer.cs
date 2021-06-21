using System;

namespace Bank.Business.Entities
{
    public class Customer
    {
        public Customer(string document, string firstName, DateTime birthDate)
        {
            Id = Guid.NewGuid();
            Document = document;
            FirstName = firstName;
            BirthDate = birthDate;
        }

        public Guid Id { get; private set; }
        public string Document { get; private set; }
        public string FirstName { get; private set; }
        public DateTime BirthDate { get; private set; }
    }
}
