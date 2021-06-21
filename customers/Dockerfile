FROM mcr.microsoft.com/dotnet/core/sdk:3.1-alpine AS build
WORKDIR /build

COPY ["customers/Bank.Api/Bank.Api.csproj", "src/Bank.Api/"]
COPY ["customers/Bank.Business/Bank.Business.csproj", "src/Bank.Business/"]

RUN dotnet restore "src/Bank.Api/Bank.Api.csproj"
COPY . .
WORKDIR /build/src/Bank.Api

FROM build AS publish
RUN dotnet publish "Bank.Api.csproj" -c Release -o /publish

FROM mcr.microsoft.com/dotnet/core/aspnet:3.1 AS app
WORKDIR /app
COPY --from=publish /publish .
ENTRYPOINT ["dotnet", "Bank.Api.dll"]