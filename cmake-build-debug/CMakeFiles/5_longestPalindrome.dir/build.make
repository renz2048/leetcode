# CMAKE generated file: DO NOT EDIT!
# Generated by "Unix Makefiles" Generator, CMake Version 3.17

# Delete rule output on recipe failure.
.DELETE_ON_ERROR:


#=============================================================================
# Special targets provided by cmake.

# Disable implicit rules so canonical targets will work.
.SUFFIXES:


# Disable VCS-based implicit rules.
% : %,v


# Disable VCS-based implicit rules.
% : RCS/%


# Disable VCS-based implicit rules.
% : RCS/%,v


# Disable VCS-based implicit rules.
% : SCCS/s.%


# Disable VCS-based implicit rules.
% : s.%


.SUFFIXES: .hpux_make_needs_suffix_list


# Command-line flag to silence nested $(MAKE).
$(VERBOSE)MAKESILENT = -s

# Suppress display of executed commands.
$(VERBOSE).SILENT:


# A target that is always out of date.
cmake_force:

.PHONY : cmake_force

#=============================================================================
# Set environment variables for the build.

# The shell in which to execute make rules.
SHELL = /bin/sh

# The CMake executable.
CMAKE_COMMAND = /home/renzheng/.local/share/JetBrains/Toolbox/apps/CLion/ch-0/203.5981.166/bin/cmake/linux/bin/cmake

# The command to remove a file.
RM = /home/renzheng/.local/share/JetBrains/Toolbox/apps/CLion/ch-0/203.5981.166/bin/cmake/linux/bin/cmake -E rm -f

# Escaping for special characters.
EQUALS = =

# The top-level source directory on which CMake was run.
CMAKE_SOURCE_DIR = /home/renzheng/CLionProjects/leetcode

# The top-level build directory on which CMake was run.
CMAKE_BINARY_DIR = /home/renzheng/CLionProjects/leetcode/cmake-build-debug

# Include any dependencies generated for this target.
include CMakeFiles/5_longestPalindrome.dir/depend.make

# Include the progress variables for this target.
include CMakeFiles/5_longestPalindrome.dir/progress.make

# Include the compile flags for this target's objects.
include CMakeFiles/5_longestPalindrome.dir/flags.make

CMakeFiles/5_longestPalindrome.dir/5-longestPalindrome/5-longestPalindrome.c.o: CMakeFiles/5_longestPalindrome.dir/flags.make
CMakeFiles/5_longestPalindrome.dir/5-longestPalindrome/5-longestPalindrome.c.o: ../5-longestPalindrome/5-longestPalindrome.c
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --progress-dir=/home/renzheng/CLionProjects/leetcode/cmake-build-debug/CMakeFiles --progress-num=$(CMAKE_PROGRESS_1) "Building C object CMakeFiles/5_longestPalindrome.dir/5-longestPalindrome/5-longestPalindrome.c.o"
	/usr/bin/cc $(C_DEFINES) $(C_INCLUDES) $(C_FLAGS) -o CMakeFiles/5_longestPalindrome.dir/5-longestPalindrome/5-longestPalindrome.c.o   -c /home/renzheng/CLionProjects/leetcode/5-longestPalindrome/5-longestPalindrome.c

CMakeFiles/5_longestPalindrome.dir/5-longestPalindrome/5-longestPalindrome.c.i: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Preprocessing C source to CMakeFiles/5_longestPalindrome.dir/5-longestPalindrome/5-longestPalindrome.c.i"
	/usr/bin/cc $(C_DEFINES) $(C_INCLUDES) $(C_FLAGS) -E /home/renzheng/CLionProjects/leetcode/5-longestPalindrome/5-longestPalindrome.c > CMakeFiles/5_longestPalindrome.dir/5-longestPalindrome/5-longestPalindrome.c.i

CMakeFiles/5_longestPalindrome.dir/5-longestPalindrome/5-longestPalindrome.c.s: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Compiling C source to assembly CMakeFiles/5_longestPalindrome.dir/5-longestPalindrome/5-longestPalindrome.c.s"
	/usr/bin/cc $(C_DEFINES) $(C_INCLUDES) $(C_FLAGS) -S /home/renzheng/CLionProjects/leetcode/5-longestPalindrome/5-longestPalindrome.c -o CMakeFiles/5_longestPalindrome.dir/5-longestPalindrome/5-longestPalindrome.c.s

CMakeFiles/5_longestPalindrome.dir/3-lengthOfLongestSubstring.c.o: CMakeFiles/5_longestPalindrome.dir/flags.make
CMakeFiles/5_longestPalindrome.dir/3-lengthOfLongestSubstring.c.o: ../3-lengthOfLongestSubstring.c
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --progress-dir=/home/renzheng/CLionProjects/leetcode/cmake-build-debug/CMakeFiles --progress-num=$(CMAKE_PROGRESS_2) "Building C object CMakeFiles/5_longestPalindrome.dir/3-lengthOfLongestSubstring.c.o"
	/usr/bin/cc $(C_DEFINES) $(C_INCLUDES) $(C_FLAGS) -o CMakeFiles/5_longestPalindrome.dir/3-lengthOfLongestSubstring.c.o   -c /home/renzheng/CLionProjects/leetcode/3-lengthOfLongestSubstring.c

CMakeFiles/5_longestPalindrome.dir/3-lengthOfLongestSubstring.c.i: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Preprocessing C source to CMakeFiles/5_longestPalindrome.dir/3-lengthOfLongestSubstring.c.i"
	/usr/bin/cc $(C_DEFINES) $(C_INCLUDES) $(C_FLAGS) -E /home/renzheng/CLionProjects/leetcode/3-lengthOfLongestSubstring.c > CMakeFiles/5_longestPalindrome.dir/3-lengthOfLongestSubstring.c.i

CMakeFiles/5_longestPalindrome.dir/3-lengthOfLongestSubstring.c.s: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Compiling C source to assembly CMakeFiles/5_longestPalindrome.dir/3-lengthOfLongestSubstring.c.s"
	/usr/bin/cc $(C_DEFINES) $(C_INCLUDES) $(C_FLAGS) -S /home/renzheng/CLionProjects/leetcode/3-lengthOfLongestSubstring.c -o CMakeFiles/5_longestPalindrome.dir/3-lengthOfLongestSubstring.c.s

# Object files for target 5_longestPalindrome
5_longestPalindrome_OBJECTS = \
"CMakeFiles/5_longestPalindrome.dir/5-longestPalindrome/5-longestPalindrome.c.o" \
"CMakeFiles/5_longestPalindrome.dir/3-lengthOfLongestSubstring.c.o"

# External object files for target 5_longestPalindrome
5_longestPalindrome_EXTERNAL_OBJECTS =

5_longestPalindrome: CMakeFiles/5_longestPalindrome.dir/5-longestPalindrome/5-longestPalindrome.c.o
5_longestPalindrome: CMakeFiles/5_longestPalindrome.dir/3-lengthOfLongestSubstring.c.o
5_longestPalindrome: CMakeFiles/5_longestPalindrome.dir/build.make
5_longestPalindrome: CMakeFiles/5_longestPalindrome.dir/link.txt
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --bold --progress-dir=/home/renzheng/CLionProjects/leetcode/cmake-build-debug/CMakeFiles --progress-num=$(CMAKE_PROGRESS_3) "Linking C executable 5_longestPalindrome"
	$(CMAKE_COMMAND) -E cmake_link_script CMakeFiles/5_longestPalindrome.dir/link.txt --verbose=$(VERBOSE)

# Rule to build all files generated by this target.
CMakeFiles/5_longestPalindrome.dir/build: 5_longestPalindrome

.PHONY : CMakeFiles/5_longestPalindrome.dir/build

CMakeFiles/5_longestPalindrome.dir/clean:
	$(CMAKE_COMMAND) -P CMakeFiles/5_longestPalindrome.dir/cmake_clean.cmake
.PHONY : CMakeFiles/5_longestPalindrome.dir/clean

CMakeFiles/5_longestPalindrome.dir/depend:
	cd /home/renzheng/CLionProjects/leetcode/cmake-build-debug && $(CMAKE_COMMAND) -E cmake_depends "Unix Makefiles" /home/renzheng/CLionProjects/leetcode /home/renzheng/CLionProjects/leetcode /home/renzheng/CLionProjects/leetcode/cmake-build-debug /home/renzheng/CLionProjects/leetcode/cmake-build-debug /home/renzheng/CLionProjects/leetcode/cmake-build-debug/CMakeFiles/5_longestPalindrome.dir/DependInfo.cmake --color=$(COLOR)
.PHONY : CMakeFiles/5_longestPalindrome.dir/depend
