# Assignment-1: Unisearcher

This project is a submission to the first assignment in PROG-2005: Cloud Technologies.  
Unisearcher is a REST web application in Golang that provides the client to retrieve information about universities 
may be candidates for application based on their name, alongside useful contextual information pertaining to the country
it is situated in.

## Endpoints

### Uniinfo

This endpoint returns information about universities that contains the given name,
and the country that the particular university is situated in.


```
Path: /unisearcher/v1/uniinfo/
Request: uniinfo/{:partial_or_complete_university_name}{?fields={:field1,field2,...}}
```
```{:partial_or_complete_university_name}``` is the partial or complete university name of the 
universities to be searched for. 

```{?fields={:field1,field2,...}}``` is an optional parameter that specifies which fields that
should be included in the result. The available fields are ```name,country,isocode,webpages,languages,map```.
If no fields are specified, **all the fields are included**.


