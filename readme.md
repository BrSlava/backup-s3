# Модуль для сбора данных 

## Сборка
Для сборки выполняем: 

```
GOOS=linux GOARCH=amd64 go build -o chk-backup cmd/main.go
```

Копируем файлы `chk-backup` и `config.toml` на ВМ , например в каталог `/etc/telegraf/telegraf.cmd`
Правим файл config.toml, предварительно настроив в asterisk нового ami пользователя
Добавляем в конфигурационный файл `chk-backup.conf` в каталоге `/etc/telegraf.d`

```
[[inputs.execd]]
command = ["/etc/telegraf/telegraf.cmd/chk-backup", "-config", "/etc/telegraf/telegraf.cmd//config.toml"]
```

## Проверка
Проверить правильность настройки можно командой

```
telegraf --test --config=/etc/telegraf/telegraf.d/chk-backup.conf --test-wait 5
```

Ответ должен сожержать что то похожее на:

```
2023-05-18T09:58:59Z I! Loading config: /etc/telegraf/telegraf.d/asterisk.conf
2023-05-18T09:58:59Z I! Starting Telegraf 1.26.2-
2023-05-18T09:58:59Z I! Available plugins: 235 inputs, 9 aggregators, 27 processors, 22 parsers, 57 outputs, 2 secret-stores
2023-05-18T09:58:59Z I! Loaded inputs: execd
2023-05-18T09:58:59Z I! Loaded aggregators:
2023-05-18T09:58:59Z I! Loaded processors:
2023-05-18T09:58:59Z I! Loaded secretstores:
2023-05-18T09:58:59Z W! Outputs are not used in testing mode!
2023-05-18T09:58:59Z I! Tags enabled: host=pbx.informunity.ru
2023-05-18T09:58:59Z I! [inputs.execd] Starting process: /opt/scripts/asterisk [-config /opt/scripts/config.toml]
> asterisk,host=pbx.informunity.ru current_call_volume=0i,last_reload=14326i,pjsip_devices_inuse=0i,pjsip_devices_notinuse=10i,pjsip_devices_total=32i,pjsip_devices_unavailable=22i,procesed_call_volume=301i,sip_monitored_offline=19i,sip_monitored_online=8i,sip_peers=28i,sip_unmonitored_offline=1i,sip_unmonitored_online=0i,system_uptime=572049i,trunc_active=6i,trunk_total=6i 1684403940862555687
> asterisk,host=pbx.informunity.ru current_call_volume=0i,last_reload=14327i,pjsip_devices_inuse=0i,pjsip_devices_notinuse=10i,pjsip_devices_total=32i,pjsip_devices_unavailable=22i,procesed_call_volume=301i,sip_monitored_offline=19i,sip_monitored_online=8i,sip_peers=28i,sip_unmonitored_offline=1i,sip_unmonitored_online=0i,system_uptime=572050i,trunc_active=6i,trunk_total=6i 1684403941863185676
> asterisk,host=pbx.informunity.ru current_call_volume=0i,last_reload=14328i,pjsip_devices_inuse=0i,pjsip_devices_notinuse=10i,pjsip_devices_total=32i,pjsip_devices_unavailable=22i,procesed_call_volume=301i,sip_monitored_offline=19i,sip_monitored_online=8i,sip_peers=28i,sip_unmonitored_offline=1i,sip_unmonitored_online=0i,system_uptime=572051i,trunc_active=6i,trunk_total=6i 1684403942861389718
> asterisk,host=pbx.informunity.ru current_call_volume=0i,last_reload=14329i,pjsip_devices_inuse=0i,pjsip_devices_notinuse=10i,pjsip_devices_total=32i,pjsip_devices_unavailable=22i,procesed_call_volume=301i,sip_monitored_offline=19i,sip_monitored_online=8i,sip_peers=28i,sip_unmonitored_offline=1i,sip_unmonitored_online=0i,system_uptime=572052i,trunc_active=6i,trunk_total=6i 1684403943860477355
2023-05-18T09:59:04Z I! [inputs.execd] Process /opt/scripts/asterisk shut down

```