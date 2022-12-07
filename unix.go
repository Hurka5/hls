//go:build !windows
// +build !windows

package main

import (
	"errors"
	"golang.org/x/sys/unix"
	"io/fs"
	"os"
	"strings"
  "os/user"
  "syscall"
  "strconv"
)

func (i Item) isHidden() bool {
	return i.info.Name()[0] == '.'
}

func (item Item) isExecutable() bool {

	m := item.info.Mode()
	/* Check if Regulat File or Symlink */
	if !((m.IsRegular()) || (uint32(m&fs.ModeSymlink) == 0)) {
		return false
	}
	/* Check if executeble  */
	if uint32(m&0111) == 0 {
		return false
	}
	/* Check if user has acces to execute it */
	if unix.Access(item.path, unix.X_OK) != nil {
		return false
	}

	return true
}

func (item Item) isLink() (bool, bool) {

	//Check if link
	isLink := uint32(item.info.Mode()&fs.ModeSymlink) != 0

	//If not link bye
	if !isLink {
		return isLink, false
	}

	//Get link target
	target, _ := os.Readlink(item.path)

	// If not absolute path make it absolute path
	dir := "/"
	if !strings.HasPrefix(target, dir) {
		dir = item.path[:strings.LastIndex(item.path, "/")+1]
	}

	isBroken := false
	if _, err := os.Stat(dir + target); errors.Is(err, os.ErrNotExist) {
		isBroken = true
	}
	return isLink, isBroken
}


func (item Item) Owner() (string, string) {
  stat := item.info.Sys().(*syscall.Stat_t)

  // Get IDs
  uid := stat.Uid
  gid := stat.Gid

  // Convert to int
  u := strconv.FormatUint(uint64(uid), 10)
  g := strconv.FormatUint(uint64(gid), 10)

  // Lookup
  usr, err := user.LookupId(u)
  check(err)
  group, err := user.LookupGroupId(g)
  check(err)

  return usr.Username, group.Name
}
