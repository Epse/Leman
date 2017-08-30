-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE Roles (
RoleID serial PRIMARY KEY,
Title text UNIQUE NOT NULL,
PermissionLevel int NOT NULL
);

CREATE TABLE Users (
       UserID serial PRIMARY KEY,
       Email text UNIQUE NOT NULL,
       FirstName text NOT NULL,
       FamilyName text NOT NULL,
       PasswordHash text NOT NULL,
       RoleID int,
       Phone text,
       CreatedOn timestamp  NOT NULL,
       LastLogin timestamp ,
       CONSTRAINT user_role_role_id_fkey FOREIGN KEY (RoleID)
                  REFERENCES Roles (RoleID) MATCH SIMPLE
                  ON UPDATE NO ACTION ON DELETE RESTRICT
);

CREATE TABLE Locations (
       LocationID serial PRIMARY KEY,
       Title text UNIQUE NOT NULL,
       IsExternal boolean NOT NULL,
       Address text,
       Phone text
);

CREATE TABLE Categories (
       CategoryID serial PRIMARY KEY,
       Title text UNIQUE NOT NULL,
);

CREATE TABLE Brands (
       BrandID serial PRIMARY KEY,
       Title text UNIQUE NOT NULL
);

CREATE TABLE Products (
       ProductID serial PRIMARY KEY,
       Title text UNIQUE NOT NULL,
       Brand int NOT NULL,
       Category int NOT NULL,
       IsIndividuallyTracked boolean NOT NULL,
       PricePerTime money NOT NULL,
       TimeUnit text NOT NULL,
       CONSTRAINT products_brand_brands_brandid_fkey FOREIGN KEY (Brand)
                  REFERENCES Brands (BrandID) MATCH SIMPLE
                  ON UPDATE NO ACTION ON DELETE RESTRICT
       CONSTRAINT products_category_categories_categoryid_fkey FOREIGN KEY (Category)
                  REFERENCES Categories (CategoryID) MATCH SIMPLE
                  ON UPDATE NO ACTION ON DELETE RESTRICT
);

CREATE TABLE Renters (
RenterID serial PRIMARY KEY,
Title text NOT NULL,
Phone text NOT NULL,
Email text NOT NULL,
Address text NOT NULL
);

CREATE TABLE Rentals (
       RentalID bigserial PRIMARY KEY,
       RentedFrom timestamp with time zone NOT NULL,
       RentedTill timestamp with time zone NOT NULL,
       CreatedOn timestamp NOT NULL,
       LastModified timestamp,
       Renter int NOT NULL,
       CONSTRAINT rentals_renter_renters_renterid_fkey FOREIGN KEY (Renter)
                  REFERENCES Renters (RenterID) MATCH SIMPLE
                  ON UPDATE NO ACTION ON DELETE RESTRICT
);

CREATE TABLE Trackables (
       TrackableID bigserial PRIMARY KEY,
       Product int NOT NULL,
       TotalQuantity int NOT NULL,
       LocationIIT int,
       QuantityAvailable int NOT NULL,
       CONSTRAINT trackables_product_products_productid_fkey FOREIGN KEY (Product)
                  REFERENCES Products (ProductID) MATCH SIMPLE
                  ON UPDATE NO ACTION ON DELETE CASCADE,
       CONSTRAINT trackables_locationiit_locations_locationid_fkey FOREIGN KEY (LocationIIT)
                  REFERENCES Locations (LocationID) MATCH SIMPLE
                  ON UPDATE NO ACTION ON DELETE RESTRICT
);

CREATE TABLE Link (
       Rental bigint NOT NULL,
       Trackable bigint NOT NULL,
       PRIMARY KEY (Rental, Trackable),
       CONSTRAINT link_rental_rentals_rentalid_fkey FOREIGN KEY (Rental)
                  REFERENCES Rentals (RentalID) MATCH SIMPLE
                  ON UPDATE NO ACTION ON DELETE CASCADE,
       CONSTRAINT link_trackable_trackables_trackableid_fkey FOREIGN KEY (Trackable)
                  REFERENCES Trackables (TrackableID) MATCH SIMPLE
                  ON UPDATE NO ACTION ON DELETE RESTRICT
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE Link;
DROP TABLE Trackables;
DROP TABLE Renters;
DROP TABLE Rentals;
DROP TABLE Products;
DROP TABLE Brands;
DROP TABLE Categories;
DROP TABLE Locations;
DROP TABLE Roles;
DROP TABLE Users;
