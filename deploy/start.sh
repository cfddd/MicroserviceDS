docker build -t router_service_image ./router_service
docker run --name router_service_container --network host -d router_service_image

docker build -t social_service_image ./social_service
docker run --name social_service_container --network host -d social_service_image

docker build -t user_service_image ./user_service
docker run --name user_service_container --network host -d user_service_image

docker build -t video_service_image ./video_service
docker run  --name video_service_container --network host -d video_service_image
