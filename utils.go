package main

import (
	"crypto/sha512"
	eb64 "encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"strconv"
	"strings"
	// "time"
)

const (
	Select = "select"
	Create = "create"
	Insert = "insert"
	Delete = "delete"
	Update = "update"
)

func checkHash(password string, hashStr string) bool {
	// now := time.Now()
	// timeFormat := now.Format("02.2006")
	hash := sha512.New()
	io.WriteString(hash, "root:"+password+":23.2019")
	
	return hashStr == eb64.RawStdEncoding.EncodeToString(hash.Sum(nil))
}

func readPasswordHash() string {
	file, err := ioutil.ReadFile("./pass.txt")

	if err != nil {
		fmt.Println("couldn't read file: " + err.Error())
		return ""
	}

	return string(file)

}

func allMake(command string)string{
	// db := map[string]*[]Student{}
	command = strings.TrimSpace(command)
		command = strings.ToLower(command)
		commandStruct := strings.Split(command, " ")

		switch commandStruct[0] {
		case Select:
			tableName := strings.TrimSpace(commandStruct[3])
			tableName=strings.TrimSpace(tableName)

			if db[tableName] == nil {
				fmt.Print(tableName)
				return fmt.Sprint("table not exits")
			} else {
				columns := commandStruct[1]
				switch columns {
				case "*":
					if len(commandStruct) > 4 {
						if commandStruct[4] == "where" {
							switch commandStruct[6] {
							case "==":

								switch commandStruct[5] {
								case "id":
									var text string
									arrayOfStudents := *db[tableName]
									text+=fmt.Sprintf("  ID|               Fname|Age|Average| \n")
									text+=fmt.Sprintln("------------------------------------")

									id_num, _ := strconv.Atoi(commandStruct[7])
									for _, row := range arrayOfStudents {

										if row.ID == id_num {

											text+=fmt.Sprintf("%4d|%20s|%3d|%2.2f\n", row.ID, row.Fname, row.Age, row.Average)

										}
									}
									return text
								}

							}
						} else {
							var text string
							text+=fmt.Sprintln(strings.Repeat("-", 50))
							text+=fmt.Sprintln("| argument of command not recognize |")
							text+=fmt.Sprintln(strings.Repeat("-", 50))
							return text
						}
					} else {
						var text string
						text+=fmt.Sprintf("  ID|               Fname|Age|Average|Experience \n")
						text+=fmt.Sprintln("------------------------------------")
						arrayOfStudents := *db[tableName]

						for _, row := range arrayOfStudents {

							text+=fmt.Sprintf("%4d|%20s|%3d|%4.2f|%2d\n", row.ID, row.Fname, row.Age, row.Average, row.Experience)
						}
						text+=fmt.Sprintln(strings.Repeat("-", 50))
						text+=fmt.Sprintln("| rows returned: ", len(arrayOfStudents))
						text+=fmt.Sprintln(strings.Repeat("-", 50))
						return text
					}
					break

				default:
					arrayOfStudents := *db[tableName]
					colls := strings.Split(commandStruct[1], ",")
					var text string
					for _, row := range arrayOfStudents {
						
						for _, coll := range colls {
							coll = strings.ToLower(coll)
							switch coll {
							case "id":
								// fmt.Printf("  ID|")
								text+=fmt.Sprintf("%4d|", row.ID)
								break
							case "fname":
								// fmt.Printf("               Fname|")
								text+=fmt.Sprintf("%20s|", row.Fname)
								break
							case "age":
								// fmt.Printf("Age|")
								text+=fmt.Sprintf("%3d|", row.Age)
								break
							case "average":
								// fmt.Printf("Average ")
								text+=fmt.Sprintf("%2.2f", row.Average)
								break
							default:
								text+=fmt.Sprintln("Ooops somefing wrong!")
							}

							// fmt.Printf("%4d|%20s|%3d|%2.2f\n", row.ID, row.Fname, row.Age, row.Average)
						}
						text+=fmt.Sprintln("")
					}
					return text
					break
				}

			}

			break
		case Create:
			tableName := commandStruct[2]
			db[tableName] = &emptySlice
			var text string
			text+=fmt.Sprintln("table created: " + tableName)
			return text
			break

		case Insert:
			var text string
			text=""
			args := commandStruct[1]
			arg := strings.Split(args, ",")
			tableName := commandStruct[3]
			age, _ := strconv.Atoi(arg[1])
			isStudent, _ := strconv.ParseBool(arg[2])
			exp, _ := strconv.Atoi(arg[3])
			if len(arg) == 4 {
				if db[tableName] == nil {
					text+=fmt.Sprintln("table not exits")
					return text
				} else {
					ID++
					emptySlice = append(emptySlice, Student{
						ID:         ID,
						Age:        age,
						Fname:      arg[0],
						IsStudent:  isStudent,
						Experience: exp,
					})
					text+=fmt.Sprint("insert one row")
				}

			} else {
				text+=fmt.Sprintln(strings.Repeat("-", 50))
				text+=fmt.Sprintln("| check value of coll or count of coll |")
				text+=fmt.Sprintln(strings.Repeat("-", 50))
				return text
			}
			db[tableName] = &emptySlice

			break

		case Delete:
			var text string
			tableName := commandStruct[1]
			emptySlice = *db[tableName]
			if commandStruct[2] == "where" {

				b, _ := strconv.Atoi(commandStruct[5])
				for _, row := range emptySlice {
					if row.ID == b {
						emptySlice = append(emptySlice[:row.ID-1], emptySlice[row.ID:]...)
					}
				}
				db[tableName] = &emptySlice

			} else {
				text+=fmt.Sprintln(strings.Repeat("-", 50))
				text+=fmt.Sprintln("| check where argument |")
				text+=fmt.Sprintln(strings.Repeat("-", 50))
				return text
			}
			break
		case Update:
			var text string
			tableName := commandStruct[4]
			arg := commandStruct[1]
			arrayOfStudents:= *db[tableName]

			b, _ := strconv.Atoi(commandStruct[8])

			if arg == "*" {
				args := commandStruct[2]
				cols := strings.Split(args, ",")
				age, _ := strconv.Atoi(cols[1])
				isStudent, _ := strconv.ParseBool(cols[2])
				exp, _ := strconv.Atoi(cols[3])
				for i, _ := range arrayOfStudents {
					if arrayOfStudents[i].ID == b {

						arrayOfStudents[i].Age = age
						arrayOfStudents[i].Fname = cols[0]
						arrayOfStudents[i].IsStudent = isStudent
						arrayOfStudents[i].Experience = exp
						
						text+=fmt.Sprintf("update all colums")
						return text
					}

				}

			}

			break
		default:
			var text string
			text+=fmt.Sprintln(strings.Repeat("-", 25))
			text+=fmt.Sprintln("| command not recognize |")
			text+=fmt.Sprintln(strings.Repeat("-", 25))
			return text
		}
		return fmt.Sprint("wait")
}