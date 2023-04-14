package main

import (
	"log"
	"time"

	"github.com/godbus/dbus/v5"
	"github.com/godbus/dbus/v5/introspect"
	"github.com/godbus/dbus/v5/prop"
)

const (
	// object path and interface
	path  = "/org/hello/Props"
	iface = "org.hello.Props"
)

func main() {
	// connect to session bus
	conn, err := dbus.ConnectSessionBus()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// request name
	reply, err := conn.RequestName(iface, dbus.NameFlagDoNotQueue)
	if err != nil {
		log.Fatal(err)
	}
	if reply != dbus.RequestNameReplyPrimaryOwner {
		// name already taken, get property
		s := ""
		err := conn.Object(iface, path).StoreProperty(iface+".Hello", &s)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(s)
		return
	}

	// properties
	propsSpec := prop.Map{
		iface: {
			"Hello": {
				Value:    "hello",
				Writable: false,
				Emit:     prop.EmitTrue,
				Callback: nil,
			},
		},
	}
	props, err := prop.Export(conn, path, propsSpec)
	if err != nil {
		log.Fatalf("export props spec failed: %v\n", err)
	}

	// introspection
	n := &introspect.Node{
		Name: path,
		Interfaces: []introspect.Interface{
			introspect.IntrospectData,
			prop.IntrospectData,
			{
				Name:       iface,
				Properties: props.Introspection(iface),
			},
		},
	}
	err = conn.Export(introspect.NewIntrospectable(n), path,
		"org.freedesktop.DBus.Introspectable")
	if err != nil {
		log.Fatalf("export introspect failed: %v\n", err)
	}

	// change properties
	go func() {
		for {
			// iterate through greetings and set Hello
			for _, hi := range []string{
				"hi",
				"hey",
				"good day",
				"hello",
			} {
				time.Sleep(5 * time.Second)
				props.SetMust(iface, "Hello", hi)
			}
		}
	}()

	// print info and keep running
	log.Println("Changing Hello property every 5 seconds")
	log.Println("View PropertiesChanged signals with: dbus-monitor --session")
	select {}
}
