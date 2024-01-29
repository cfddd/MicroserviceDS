# 在MicroserviceDS项目根路径下运行docker命令
docker build -t router_service_image ./router_service
docker run -v ./router_service/configs:/build/configs --name router_service_container --network host -d router_service_image

docker build -t social_service_image ./social_service
docker run -v ./social_service/config:/build/config --name social_service_container --network host -d social_service_image

docker build -t user_service_image ./user_service
docker run -v ./user_service/config:/build/config --name user_service_container --network host -d user_service_image

docker build -t video_service_image ./video_service
docker run  -v ./video_service/config:/build/config --name video_service_container --network host -d video_service_image

# 还需要配置ffmpeg，参考https://zhuanlan.zhihu.com/p/347780238