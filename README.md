![Logo](/assets/portadareadme.webp)

# SSH-COMANDOCKER

Este proyecto nace de una necesidad técnica real dentro de mi entorno laboral, donde gestionamos una infraestructura compuesta por un servidor principal y más de 15 servidores distribuidos. Lo curioso de este escenario es que, aunque cada servidor externo es idéntico en estructura, cada uno aloja una instancia de la misma aplicación que procesa información distinta dependiendo de la localidad donde el servidor esté físicamente instalado.

### El Desafío de la Infraestructura
Mantener esta red de servidores se volvió una tarea extremadamente tediosa, especialmente al momento de actualizar las aplicaciones o gestionar los contenedores de Docker. Ejecutar procesos de mantenimiento de forma secuencial —entrando servidor por servidor— consumía un tiempo valioso y aumentaba la complejidad operativa. Ante este problema, decidí desarrollar esta referencia de proyecto, buscando una solución basada en **código limpio** y eficiencia técnica sin necesidad de recurrir a herramientas de orquestación excesivamente pesadas.

### La Solución: Concurrencia en Go
**SSH-Comandocker** es una mini aplicación escrita en Go que ofrece una solución rápida y potente para la conexión masiva con estos servidores externos. El funcionamiento es directo: en lugar de procesar cada servidor de uno en uno, la aplicación utiliza las capacidades de concurrencia de Go para crear clientes SSH simultáneos. 

Al ejecutar la herramienta con las conexiones y credenciales definidas, el sistema inicia un proceso en paralelo. Esto significa que, mientras se establece la comunicación con un servidor, se está haciendo lo mismo con todos los demás al mismo tiempo. Es un proceso maravilloso que permite ejecutar una serie de comandos o actualizaciones críticas en toda la infraestructura de manera inmediata.

### Desarrollo completo
El desarrollo más robusto y sofisticado que realicé para la empresa se fundamentó en una arquitectura de gestión centralizada, diseñada para manejar con precisión la lógica de conexión de múltiples nodos. En lugar de configuraciones locales, el sistema dependía de una base de datos central donde se alojaban de forma segura las credenciales y metadatos de clientes y servidores. Al iniciar la ejecución, el programa orquestaba una recopilación dinámica de todas las conexiones activas, permitiendo desplegar tareas masivas o comandos específicos de manera simultánea en todo el parque de servidores registrados.

Para garantizar la fiabilidad del proceso, el sistema contaba con un motor de trazabilidad avanzado que separaba los logs por hilos de ejecución. Esto permitía un monitoreo detallado paso a paso, facilitando la identificación inmediata de qué cliente presentaba un error y cuál finalizaba con éxito, generando así un respaldo íntegro de cada operación. Además, integré funciones nativas de Linux para la manipulación del entorno, desde la verificación de estados del sistema hasta la gestión de variables y archivos. Esta capa de control no solo automatizaba tareas, sino que actuaba como un filtro de seguridad proactivo que analizaba los comandos introducidos por el usuario para restringir acciones potencialmente peligrosas.


**Si decides evolucionar este proyecto, el potencial de escalabilidad es inmenso. Podrías rediseñar el módulo de autenticación para implementar métodos más modernos, o incluso transformar la interfaz de consola en un panel de control dinámico que muestre el flujo de datos en vivo. La estructura actual es una base sólida y simple que invita a integrar nuevas funciones predefinidas o permitir una interacción más profunda mediante una consola de comandos enviada directamente por el usuario, manteniendo siempre el control centralizado que define a esta herramienta.**


### Notas sobre el Desarrollo
Es importante notar que esta aplicación refleja la lógica y la solución que implementé en mi empresa, por lo que el código puede presentarse como una referencia técnica de la arquitectura de solución más que como un producto comercial acabado. (Un codigo mas complejo y estructurado esta 'Activo' y desarrollado en la empresa al cual realice el darrollo)

**Consideración de Seguridad:**
Actualmente, el código utiliza una autenticación "insegura" basada en usuario y contraseña a modo de ejemplo y facilidad de uso inicial. En un entorno de producción real, lo ideal sería evolucionar esta conexión hacia el uso de certificados o llaves SSH para garantizar un estándar de seguridad mucho más elevado.