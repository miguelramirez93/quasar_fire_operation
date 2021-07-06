  
docker_compose(['./docker-compose.yml'])
docker_build('local/quasar_fire_operation_api', '.', dockerfile = 'Dockerfile',
    live_update=[
        sync('.', '/go/src/github.com/miguelramirez93/quasar_fire_operation'),
    ],
    target="dev"
)