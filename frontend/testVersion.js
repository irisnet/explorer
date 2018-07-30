export default {
  app:{
    "url":"https://bianjie.ai/rainbow/${env}/rainbow-latest.apk",

    "version": "${app.version}",

    "force_update": false,

    "env":"${env}",

    "buildNumber":"${bamboo.buildNumber}",

    "apiServerIP":"${server.ip}",

    "apiServerPort":"${server.port}",

    "explorer":"${explorer}",

    "chainId":"${chainId}",
  }
}
