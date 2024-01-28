docker build -t router_service_image ./router_service
docker run -d --network host --name router_service_container

docker build -t social_service_image ./social_service
docker run -d --network host --name social_service_container

docker build -t user_service_image ./user_service
docker run -d --network host --name user_service_container

docker build -t video_service_image ./video_service
docker run -d --network host --name video_service_container
