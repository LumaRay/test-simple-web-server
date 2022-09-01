// https://github.com/drogonframework/drogon/wiki/ENG-03-Quick-Start

#include <drogon/drogon.h>

int main() {
    //Set HTTP listener address and port
    //drogon::app().addListener("0.0.0.0",80);
    drogon::app().addListener("127.0.0.1",8080);
    drogon::app().setThreadNum(4);
    //drogon::app().setLogLevel(trantor::Logger::LogLevel::kInfo);
    //TRACE, DEBUG, INFO, WARN
    //Load config file
    //drogon::app().loadConfigFile("../config.json");
    //Run HTTP framework,the method will block in the internal event loop
    drogon::app().run();
    return 0;
}
