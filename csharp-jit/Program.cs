using Microsoft.AspNetCore.Builder;
using Microsoft.AspNetCore.Hosting;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.Hosting;
using Npgsql;
using Dapper;

var builder = WebApplication.CreateBuilder(args);

// Add configuration for PostgreSQL connection string
var connectionString = builder.Configuration.GetConnectionString("PostgresConnection");

var app = builder.Build();

// Endpoint to fetch data from PostgreSQL
app.MapGet("/users", async () =>
{
    using var connection = new NpgsqlConnection(connectionString);
    var users = await connection.QueryAsync<User>("SELECT id, name, email FROM users");
    return Results.Ok(users);
});

// Default endpoint
app.MapGet("/", () => "Hello, World!");

app.Run();

// Define a simple User model
public record User(int Id, string Name, string Email);