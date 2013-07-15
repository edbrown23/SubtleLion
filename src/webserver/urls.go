package webserver

import (
    "regexp"
)

/*
Manages url regular expressions and their associated callback functions
for the webserver. Direct access is hidden from the users of the webserver
*/

type regexpCallback struct {
    pattern *regexp.Regexp
    callback interface{}
}


type urls struct {
    patterns map[string]regexpCallback
}

// Should not be exposed
func Newurls() *urls {
    urls := new(urls)
    urls.patterns = make(map[string]regexpCallback)
    return urls
}

// Should be exposed, not tied to urls, probably to Webserver
func (urls *urls) RegisterCallback(url string, callback interface{}) error {
    rexp, error := regexp.Compile(url)
    if (error != nil) {
        return error
    }
    rexpcall := regexpCallback{rexp, callback}

    urls.patterns[url] = rexpcall

    return nil
}