package server

const mascot = `

                            ,▄▄▓▓████████████▓▓▄▄,
               ▄▄mmæ▄████▓██╬╬╬╬╬╠╠╠╠╠╠╬╬╬╬╬╬╬╬╬╬╬▓██▄▄
            ▄████   ▄██╬▒╬╬╬╬╬╬╬╬╬╠╠╠╠╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬██▄,
          ,███▀└ ▄██╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬████▓▄,
          █    ▄█╬▒╠╠╠╠╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╠╠╬╬██▓▄
          █  ▓█▒╠╬╠╠╠╠╬╬╬╬╬╬╬╬╬╬▒▓▓▓▓▓▓▓╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╣▓▓▓▓╬╬╬╬╠▒╬██▄ç
         ▓███▒╬╬╠╠╠╠╠╬╬╬╬╬╬╬╬▓█▀╬╬╬╬╬╬████╬▓╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╣▓▓▓▓▓╬╬╬╬╠╠╬▀██▄
         ╟█▒╠╬╬╠╠╠╬╬╬╬╬╬╬╬╬╬╬▄▓██▀▀▀▀▀█▓▒╬██╬▓╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╣▓▓▓▓▓▓╬╬╬╬╬╬╬╬╬█▓
        ▓█▒╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬▓█╙ ╓▄▄▄    ╙▀█▒╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╣▓▓▓▓▓▓▓▓╬╣╬╬╬╣╬╬██µ
       ██╠╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬█▀ █▀█▀╙█▌     ╙█▒╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬▓▓▓▓▓▓▓▓▓▓╬╬╬▓▓▓╬╬██
      ██╬╬╬╠╬╬╬╬╬╬╬╬╬╬╬╬╬╣█ █▌██▄██       ▐█╬╬╠╠╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╣▓▓▓▓▓▓█╬▓▓▓▓▓▓▓▓▓▓╣█
     ▐█╬╣▓▒╠╬╬╬╬╬╬╬╬╬╬╬╬╬╣█ └▀██▀─         █╬╬╠╠╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╣▓▓▓▓▓▓▓█▓▓╬█████████
     █▒███▌╬╬╬╬╬╬╬╬╬╬╬╬╬╠╠█▄              ▐█╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬▓▓▓▓▓▓▓╬█▀╙└   └└─
    ▐██╬█▌╬╬╬╬╬╬╬╬╬╬╬╬╬╬╠╠░█▌           ,██╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╣▓▓▓▓▓▓▓▓█▄
    █████╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╠╬▒╬██▄▄,,,▄▄▓██░╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╣╬╬╣▓▓▓▓▓▓▓▓█▌
    █▒▓██╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬██▓▒╬╬╬╬╬╬╬▒▒███▓▓╬╬╬╬╬╬╬╬╬╬╬╬╬╬╣▓▓▒╣▓▓▓▓▓▓▓▓╣█
    █▌╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬██████████╬▓▓▓╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬▓▓▒╬▓▓▓▓▓▓▓▓╣█
    ██╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╣▓▓▓╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╣█╬▓▓▓▓▓██
    ╟█╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬██╬▓▓▓█▌
     ██╬╬╬╬╬╬╬╬╬╠╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬███╣█─
     ╙█▓╬╬╬╬╬╬╬╬╠╠╠╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬██
      ██▓╬╬╬╬╬╬╬╬╠╠╠╠╠╠╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬█▄
       ██▓╬╬╬╬╬╬╬╬╬╠╠╠╠╠╠╠╬╬╬╠╬╠╠╠╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╣▓▓╬╬╣█▄
        ▀█╬╬╬╬╬╬╬╬╬╬╬╬╠╠╠╠╬╬╬╠╠╠╠╠╠╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╣▓▓▓▓╬╬██
         ╙██▓╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╠╠╠╠╠╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╣▓▓▓▓▓▓▓▓▓╣╬╬╬╬╣╣╬╬╬▓╬╬╣╬█▌
           ▀██▓╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬▓▓▓╬╬╬███████▓▓╣╬╬╬╬╣╣▓▓▓██
             ▀██╬▓╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬▓▌╫▓▓▓▓▓▓╬██╙╙▀▀▀███████▀▀
               ╙██▓▓▓╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╣▓▓▓▓▓██▀─
                 └▀███╬▓▓▓╣╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╬╣▓▓▓╬██▀▀
                    ─╙▀███╬▓▓▓▓▓▓▓▓▓▓▓╬╬╬╬╬╬╬╬╬╣▓▓▓╬████¬
                         ╙▀█████▓╬▓▓▓▓╬╬╬╬╣▓▓▓▓▓╫███╬▓╬█µ
                               ─╙██╬▓▓▓╬╬╬╬╬╬╬╬╬╬╬╬╬▓▓▓╬█▄
                                   ╙▀██▓▓▓▓▓▓▓▓▓▓╬╬╬╬╬╣▓▓██
                                       ╙▀▀████▓╬╬▒╬╬╬╬╬▓███
                                             └╙╙▀▀▀▀▀▀▀▀╙`