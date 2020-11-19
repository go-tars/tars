package template

var (
	Config = `<tars>
        <application>
                <server>
                        app={{.App}}
                        server={{.Server}}
                        local=tcp -h 127.0.0.1 -p 10014 -t 30000
			logpath=/tmp
                        <{{.App}}.{{.Server}}.{{.Servant}}ObjAdapter>
                                allow
                                endpoint=tcp -h 127.0.0.1 -p 10015 -t 60000
                                handlegroup={{.App}}.{{.Server}}.{{.Servant}}ObjAdapter
                                maxconns=200000
                                protocol=tars
                                queuecap=10000
                                queuetimeout=60000
                                servant={{.App}}.{{.Server}}.{{.Servant}}Obj
                                shmcap=0
                                shmkey=0
                                threads=1
                        </{{.App}}.{{.Server}}.{{.Servant}}ObjAdapter>
                </server>
        </application>
</tars>
`
)
