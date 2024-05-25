# Modelos que corresponden a la base de datos de Northwind

* **Customers** (Clientes): Información sobre los clientes, incluyendo ID, nombre de la empresa, contacto, dirección, ciudad, región, código postal, país, teléfono y fax.

* **Employees** (Empleados): Detalles sobre los empleados de la empresa, como ID del empleado, nombre, título, título de cortesía, fecha de nacimiento, fecha de contratación, dirección, ciudad, región, código postal, país, teléfono, extensión, notas y un campo de imagen.

* **Orders** (Pedidos): Información sobre los pedidos realizados por los clientes, incluyendo ID del pedido, ID del cliente, ID del empleado, fecha del pedido, fecha de envío, fecha requerida, ID del transportista, dirección de envío, ciudad de envío, región de envío, código postal de envío, país de envío y carga.

* **Order Details** (Detalles del pedido): Detalles específicos de cada pedido, como ID del pedido, ID del producto, precio unitario, cantidad, descuento.

* **Products** (Productos): Datos sobre los productos disponibles para la venta, incluyendo ID del producto, nombre del producto, ID del proveedor, ID de la categoría, cantidad por unidad, precio unitario, unidades en stock, unidades en pedido, nivel de pedido y si el producto está discontinuado.

* **Suppliers** (Proveedores): Información sobre los proveedores, tales como ID del proveedor, nombre de la empresa, nombre de contacto, título de contacto, dirección, ciudad, región, código postal, país, teléfono, fax, página web.

* **Shippers** (Transportistas): Detalles sobre las compañías de transporte, incluyendo ID del transportista, nombre de la compañía y teléfono.

* **Categories** (Categorías): Información sobre las categorías de productos, como ID de la categoría, nombre de la categoría y descripción.

* **Region** (Región): Datos sobre las regiones, incluyendo ID de la región y descripción de la región.

* **Territories** (Territorios): Información sobre los territorios de ventas, incluyendo ID del territorio, descripción del territorio y ID de la región.

* **EmployeeTerritories** (Territorios de empleados): Relación entre empleados y territorios, con ID de empleado e ID de territorio.