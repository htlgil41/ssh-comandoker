![Logo](/assets/portadareadme.webp)

# SSH-COMANDOCKER

Este proyecto nace de una necesidad técnica real dentro de mi entorno laboral, donde gestionamos una infraestructura compuesta por un servidor principal y más de 15 servidores distribuidos. Lo curioso de este escenario es que, aunque cada servidor externo es idéntico en estructura, cada uno aloja una instancia de la misma aplicación que procesa información distinta dependiendo de la localidad donde el servidor esté físicamente instalado.

### El Desafío de la Infraestructura
Mantener esta red de servidores se volvió una tarea extremadamente tediosa, especialmente al momento de actualizar las aplicaciones o gestionar los contenedores de Docker. Ejecutar procesos de mantenimiento de forma secuencial —entrando servidor por servidor— consumía un tiempo valioso y aumentaba la complejidad operativa. Ante este problema, decidí desarrollar esta referencia de proyecto, buscando una solución basada en **código limpio** y eficiencia técnica sin necesidad de recurrir a herramientas de orquestación excesivamente pesadas.

### La Solución: Concurrencia en Go
**SSH-Comandocker** es una mini aplicación escrita en Go que ofrece una solución rápida y potente para la conexión masiva con estos servidores externos. El funcionamiento es directo: en lugar de procesar cada servidor de uno en uno, la aplicación utiliza las capacidades de concurrencia de Go para crear clientes SSH simultáneos. 

Al ejecutar la herramienta con las conexiones y credenciales definidas, el sistema inicia un proceso en paralelo. Esto significa que, mientras se establece la comunicación con un servidor, se está haciendo lo mismo con todos los demás al mismo tiempo. Es un proceso maravilloso que permite ejecutar una serie de comandos o actualizaciones críticas en toda la infraestructura de manera inmediata.


### Notas sobre el Desarrollo
Es importante notar que esta aplicación refleja la lógica y la solución que implementé en mi empresa, por lo que el código puede presentarse como una referencia técnica de la arquitectura de solución más que como un producto comercial acabado. (Un codigo mas complejo y estructurado esta 'Activo' y desarrollado en la empresa al cual realice el darrollo)

**Consideración de Seguridad:**
Actualmente, el código utiliza una autenticación "insegura" basada en usuario y contraseña a modo de ejemplo y facilidad de uso inicial. En un entorno de producción real, lo ideal sería evolucionar esta conexión hacia el uso de certificados o llaves SSH para garantizar un estándar de seguridad mucho más elevado.