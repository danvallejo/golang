using System;
using Amazon;
using Amazon.DynamoDBv2;
using Amazon.DynamoDBv2.DataModel;
using Amazon.Runtime;

namespace DynamoDb
{
    [DynamoDBTable("Students")]
    public class Student
    {
        [DynamoDBHashKey]
        public int Id { get; set; } 

        public string Name { get; set; }
    }

    class Program
    {

        static void Main(string[] args)
        {
            try
            {
                var creds = new ;

                var client = new AmazonDynamoDBClient(creds, RegionEndpoint.USWest2);

                var context = new DynamoDBContext(client);

                var student = context.Load<Student>(101);

            }
            catch (Exception ex)
            {
                Console.Write(ex.Message);

            }

            Console.Write("Done.");
            Console.ReadLine();
        }
    }
}
