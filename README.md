[![Build Status](https://travis-ci.com/miguelramirez93/quasar_fire_operation.svg?branch=master)](https://travis-ci.com/miguelramirez93/quasar_fire_operation)

# quasar_fire_operation

go libs used:
- [gin](https://github.com/gin-gonic/gin)
- [testify](https://github.com/stretchr/testify)

To deploy localy:
1. download and install [tilt](https://docs.tilt.dev/install.html)
2. Install docker and docker-compose
3. Run ```tilt up --hud=true``` at the root of this project

NOTE: Dockefile will run test coverage for you at build step ;)

After your environment is ready, you can navigate to:
- http://localhost:8080/swagger/index.html to see swagger documentation


current deployed env:
- https://localhost:8080/swagger/index.html