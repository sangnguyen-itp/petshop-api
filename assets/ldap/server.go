package ldap

import (
	"github.com/vjeantet/ldapserver"
	"log"
	"os"
	"os/signal"
	"syscall"
)

type LDAPServer interface {
	Run(addr string)
}

type ldapServer struct {
	server *ldapserver.Server
}

func (l *ldapServer) Run(addr string) {
	// listen on 10389
	go l.server.ListenAndServe(addr)

	// When CTRL+C, SIGINT and SIGTERM signal occurs
	// Then stop server gracefully
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
	close(ch)

	l.server.Stop()
}

func NewServer() LDAPServer {
	server := ldapserver.NewServer()
	routes := newRoute()
	routes.Bind(handleBind)
	routes.Search(handleSearch)
	routes.Add(handleAdd)
	server.Handle(routes)

	return &ldapServer{
		server: server,
	}
}

func newRoute() *ldapserver.RouteMux {
	return ldapserver.NewRouteMux()
}

func handleBind(w ldapserver.ResponseWriter, m *ldapserver.Message) {
	br := m.GetBindRequest()
	res := ldapserver.NewBindResponse(ldapserver.LDAPResultSuccess)
	log.Printf("Bind success User=%s, Pass=%s", string(br.Name()), string(br.AuthenticationSimple()))
	w.Write(res)
}

func handleSearch(w ldapserver.ResponseWriter, m *ldapserver.Message) {
	sr := m.GetSearchRequest()
	res := ldapserver.NewSearchResultEntry(sr.FilterString())
	log.Println(sr)
	w.Write(res)
}

func handleAdd(w ldapserver.ResponseWriter, m *ldapserver.Message) {
	ar := m.GetAddRequest()
	res := ldapserver.NewAddResponse(ldapserver.)
	log.Println(ar)
	w.Write(res)
}